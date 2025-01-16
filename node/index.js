import express from 'express';
import { createServer } from 'node:http';
import { Server } from 'socket.io';
import { v4 as uuidv4 } from 'uuid';

const app = express();
const server = createServer(app);
const io = new Server(server, {
    cors: {
        origin: "http://localhost",
    }
});

// key - socket id of board
// value - room id
const activeRooms = new Map();

io.on('connection', (socket) => {

    socket.on('boardJoinRoom', () => { 
        const roomID = uuidv4();
        console.log('Board joined room: ' + roomID);
        socket.join(roomID);
        activeRooms.set(socket.id, roomID);
        io.emit('roomsUpdated', JSON.stringify(Object.fromEntries(activeRooms)));
    });

    socket.on('controlJoinRoom', (roomID) => {
        console.log('Socket joined room - ' + roomID);
        socket.join(roomID);
    });

    socket.on('disconnect', () => {
        console.log('disconnected');
        console.log(socket.id);
        activeRooms.delete(socket.id);
        io.emit('roomsUpdated', JSON.stringify(Object.fromEntries(activeRooms)));
    });

});

app.get('/api/rooms', (req, res) => {
    res.json(Object.fromEntries(activeRooms));
});

server.listen(3000, () => {
    console.log('server running at http://localhost:3000');
});