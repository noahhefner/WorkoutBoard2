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

// key - room id
// value - socket id of board
const activeRooms = new Map();

io.on('connection', (socket) => {

    socket.on('boardJoinRoom', () => { 
        const roomID = uuidv4();
        socket.join(roomID);
        activeRooms.set(roomID, socket.id);
        socket.emit('boardIDCreated', roomID);
        io.emit('roomsUpdated', JSON.stringify(Object.fromEntries(activeRooms)));
    });

    socket.on('controlJoinRoom', (roomID) => {
        socket.join(roomID);
        io.to(roomID).emit('controlConnected');
    });

    socket.on('disconnect', () => {

        for (let [key, value] of activeRooms) {
            if (value === socket.id) {
                activeRooms.delete(key);
            }
        }

        io.emit('roomsUpdated', JSON.stringify(Object.fromEntries(activeRooms)));
    });

    socket.on('workoutSelected', (workoutID) => {
        socket.rooms.forEach((roomID) => {
            io.to(roomID).emit('workoutSelected', workoutID);
        });
    });

});

app.get('/api/rooms', (req, res) => {
    console.log(activeRooms);
    res.json(Object.fromEntries(activeRooms));
});

server.listen(3000, () => {
    console.log('server running at http://localhost:3000');
});