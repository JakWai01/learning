const peerConnection = new RTCPeerConnection;
peerConnection.ondatachannel = event => {
    var channel = event.channel;
    
    channel.onopen = event => {
        channel.send("A small step for you a big step for me");
    }

    channel.onmessage = event => {
        console.log(event.data);
    }
}
