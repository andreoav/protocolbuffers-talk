const fs = require('fs');
const userpb = require('./types/js/user_pb');

const fileBuffer = fs.readFileSync('./data/user.bin');
const user = userpb.User.deserializeBinary(fileBuffer);

console.log(user.toObject());
