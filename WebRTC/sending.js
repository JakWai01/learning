const peerConnection = new RTCPeerConnection();
// ordered can be set to false if the order of the messages does not matter
// maxPacketLifeTime The maximum number of milliseconds that attempts to transfer a message may take in unreliable mode
// maxRetransmits The maximum number of times the user agent should attempt to retransmit a message which fails the first time in unreliable mode
// protocol The name of the sub-protocol being used on the RTCDataChannel, if any.
// negotiated true if you want to be able to call createDataChannel from both sides, false if only from one
// id An 16-bit numeric ID for the channel; permitted values are 0-65534. If you don't include this option, the user agent will select an ID for you. 
const dataChannel = peerConnection.createDataChannel("testChannel", [
  (ordered = true),
  (maxPacketLifeTime = null),
  (maxRetransmits = null),
  (protocol = ""),
  (negotiated = false),
  (id = null),
]);

dataChannel.onopen = event => {
    channel.send("Hello World!")    
};

dataChannel.onmessage = event => {
    console.log(event.data)
};
