const express = require('express')
const app = express()
const server = require('http').Server(app)
const io = require('socket.io')(server)
const { v4: uuidV4 } = require('uuid')

app.set('view engine', 'ejs')
app.use(express.static('public'))

// It's a peer to peer connection, so everyone can send whatever they want to. Therefore we need correction mechanisms or control over more that 50% of the machines.
// someone who connects new to the site will get a unique roomid
app.get('/', (req, res) => {
    res.redirect(`/${uuidV4()}`)
})

// declaring said userId and rendering the view "room"
app.get('/:room', (req, res) => {
    res.render('room', { roomId: req.params.room })
})

io.on('connection', socket => {
// listening to 'join-room'. Once we receive something, execute the function 
    socket.on('join-room', (roomId, userId) => {
        socket.join(roomId)
        socket.to(roomId).broadcast.emit('user-connected', userId)
    })
})

server.listen(3000)