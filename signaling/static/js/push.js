'use strict';

const localVideo = document.getElementById('localVideo');
const remoteVideo = document.getElementById('remoteVideo');

const startPushBtn = document.getElementById('btnStartPush');
const stopPushBtn = document.getElementById('btnStopPush');
const startPullBtn = document.getElementById('btnStartPull');
const stopPullBtn = document.getElementById('btnStopPull');


startPushBtn.addEventListener('click', startPush);
stopPushBtn.addEventListener('click', stopPush);
startPullBtn.addEventListener('click', startPull);
stopPullBtn.addEventListener('click', stopPull);

var uid = $("#uid").val();
var streamName = $("#streamName").val();
var audio = $("#audio").val();
var video = $("#video").val();

const config = {};

const offerOptions = {
	offerToReceiveAudio: false,
	offerToReceiveVideo: false
};

let pc1 = new RTCPeerConnection(config);
let pc2 = new RTCPeerConnection(config);
let remoteStream;

function startPush() {
	console.log("start push: /signaling/push");

	$.post("/signaling/push",
		{"uid": uid, "streamName": streamName, "audio": audio, "video": video},
		function(data, textStatus){
		},
		"json"
	)

	var constraints = {
		audio: false,
		video: true,
	};

	navigator.mediaDevices.getDisplayMedia(constraints)
		.then(handleSuccess).catch(handleError);

}

function stopPush() {
	console.log("pc1 stop push stream");

	if (pc1) {
		pc1.close();
		pc1 = null;
	}

	localVideo.srcObject = null;
}

function startPull() {
	console.log("start pull stream");
	
	remoteVideo.srcObject = remoteStream;

	pc2.createAnswer().then(
		onCreateAnswerSuccess,
		onCreateSessionDescriptionError
	);

}

function stopPull() {
        console.log("pc2 stop pull stream");

	if (pc2) {
		pc2.close();
		pc2 = null;
	}

	remoteVideo.srcObject = null;
}


function onCreateAnswerSuccess(desc) {
	console.log('answer from pc2: \n' + desc.sdp);

	console.log('pc2 set local description start');
	pc2.setLocalDescription(desc).then(
		function() {
			onSetLocalSuccess(pc2);
		},
		onSetSessionDescriptionError
	);

	//sdp exchange
	pc1.setRemoteDescription(desc).then(
		function() {
			onSetRemoteSuccess(pc1);
		},
		onSetSessionDescriptionError
	);
}

function handleSuccess(stream) {
	console.log("get screen stream success");
	localVideo.srcObject = stream;

	pc1.oniceconnectionstatechange = function(e) {
		onIceStateChange(pc1, e);
	};

	pc1.onicecandidate = function(e) {
		onIceCandidate(pc1, e)
	}

	pc1.addStream(stream);

	pc1.createOffer(offerOptions).then(
		onCreateOfferSuccess,
		onCreateSessionDescriptionError
	);
}

function getPc(pc) {
	return pc == pc1 ? 'pc1' : 'pc2';
}

function onCreateOfferSuccess(desc) {
	console.log('offer from pc1: \n' + desc.sdp);


	console.log('pc1 set local description start');
	pc1.setLocalDescription(desc).then(
		function() {
			onSetLocalSuccess(pc1);
		},
		onSetSessionDescriptionError
	);

	//sdp exchange
	pc2.oniceconnectionstatechange = function(e) {
		onIceStateChange(pc2, e);
	}

	pc2.onicecandidate = function(e) {
		onIceCandidate(pc2, e);
	}

	pc2.onaddstream = function(e) {
		console.log('pc2 receive stream, stream_id: ' + e.stream.id);
		remoteStream = e.stream;
		//	remoteVideo.srcObject = e.stream;
	}

	pc2.setRemoteDescription(desc).then(
		function() {
			onSetRemoteSuccess(pc2);
		},
		onSetSessionDescriptionError
	);
}

function onSetLocalSuccess(pc) {
	console.log(getPc(pc) + ' set local success');
}

function onSetRemoteSuccess(pc) {
	console.log(getPc(pc) + ' set remote success');
}

function onCreateSessionDescriptionError(error){
	console.log('create session description error: ' + error.toString());
}

function onSetSessionDescriptionError(error) {
	console.log('set session description error: ' + error.toString());
}

function onIceStateChange(pc, e) {
	console.log(getPc(pc) + ' ice state change: ' + pc.iceConnectionState);
}

function getOther(pc) {
	return pc == pc1 ? pc2 : pc1;
}

function onIceCandidate(pc, e) {
	console.log(getPc(pc) + ' get new ice candidate: ' + 
		(e.candidate ? e.candidate.candidate : '(null)'));

	getOther(pc).addIceCandidate(e.candidate).then(
		function() {
			console.log(getPc(getOther(pc)) + ' add ice candidate success');
		},
		function(err) {
			console.log(getPc(getOther(pc)) + ' add ice candidate error: ' + err.toString());
		}
	);
}

function handleError(err) {
	console.log("get screen stream error: ", err.toString());
}


