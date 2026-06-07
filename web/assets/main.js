const wsProtocols = window.location.protocol == "http:" ? "ws:" : "wss:";
const ws = new WebSocket(
    wsProtocols + "//" + window.location.host + "/loket/ws",
);
let data = {};
let refresh = () => {};

ws.onmessage = (event) => {
    data = JSON.parse(event.data);
    refresh();
};

function nextAntrian(loket) {
    ws.send(loket);
}
