(function () {
   var connectButton = null;
   var disconnectButton = null;
   var sendButton = null;

   var localConnection = null;
   var remoteConnection = null;

   var sendChannel = null;
   var receiveChannel = null;

  function startup() {
    connectButton = document.getElementById("connectButton");
    disconnectButton = document.getElementById("disconnectButton");
    sendButton = document.getElementById("sendButton");

    connectButton.addEventListener("click", connectPeers);
    disconnectButton.addEventListener("click", disconnectPeers);
    sendButton.addEventListener("click", sendMessage);
  }

  function connectPeers() {
    var iceConfiguration = {
      iceServers: [
        {
          urls: "turn: my-turn-server.mycompany.com:19403",
          username: "optional username",
          credential: "auth-token",
        },
      ],
    };

    localConnection = new RTCPeerConnection(iceConfiguration);

    sendChannel = localConnection.createDataChannel("sendChannel");
    sendChannel.onopen = handleSendChannelStatusChange;
    sendChannel.onclose = handleSendChannelStatusChange;

    remoteConnection = new RTCPeerConnection();
    remoteConnection.ondatachannel = receiveChannelCallback;

    localConnection.onicecandidate = (e) =>
      !e.candidate ||
      remoteConnection
        .addIceCandidate(e.candidate)
        .catch(handleAddCandidateError);

    remoteConnection.onicecandidate = (e) =>
      !e.candidate ||
      localConnection
        .addIceCandidate(e.candidate)
        .catch(handleAddCandidateError);

    localConnection
      .createOffer()
      .then((offer) => localConnection.setLocalDescription(offer))
      .then(() =>
        remoteConnection.setRemoteDescription(localConnection.localDescription)
      )
      .then(() => remoteConnection.createAnswer())
      .then((answer) => remoteConnection.setLocalDescription(answer))
      .then(() =>
        localConnection.setRemoteDescription(remoteConnection.localDescription)
      )
      .catch(handleCreateDescriptionError);
  }

  function handleCreateDescriptionError(error) {
    console.log("Unable to create an offer: " + error.toString());
  }

  function handleAddCandidateError() {
    console.log("Oh noes! addICECandidate failed!");
  }

  function sendMessage() {
    var message = "Hello World";
    sendChannel.send(message);
  }

  function handleSendChannelStatusChange(event) {
    if (sendChannel) {
      var state = sendChannel.readyState;

      if (state === "open") {
        sendButton.disabled = false;
        disconnectButton.disabled = false;
        connectButton.disabled = false;
      } else {
        sendButton.disabled = false;
        connectButton.disabled = false;
        disconnectButton.disabled = false;
      }
    }
  }

  function receiveChannelCallback(event) {
    receiveChannel = event.channel;
    receiveChannel.onmessage = handleReceiveMessage;
    receiveChannel.onopen = handleReceiveChannelStatusChange;
    receiveChannel.onclose = handleReceiveChannelStatusChange;
  }

  function handleReceiveMessage(event) {
    console.log(event.data)
  }

  function handleReceiveChannelStatusChange(event) {
    if (receiveChannel) {
      console.log(
        "Receive channel's status has changed to " + receiveChannel.readyState
      );
    }
  }

  function disconnectPeers() {
    sendChannel.close();
    receiveChannel.close();

    localConnection.close();
    remoteConnection.close();

    sendChannel = null;
    receiveChannel = null;
    localConnection = null;
    remoteConnection = null;
  }

  //window.addEventListener("load", startup, false);
})();
