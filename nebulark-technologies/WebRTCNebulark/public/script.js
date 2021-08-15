const socket = io('/')

// You can also run your own PeerServer if you don't like the cloud 
const myPeer = new Peer(undefined, {
    host: '/',
    port: '3001'
})

// Every Peer object is assigned a unique,random ID when it is created
myPeer.on('open', id => {
    socket.emit('join-room', ROOM_ID, id)
})

socket.on('user-connected', userId => {
    console.log('Connection established with user: ' + userId)
})