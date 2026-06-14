const wsProtocols = window.location.protocol == "http:" ? "ws:" : "wss:";
const host = import.meta.env.VITE_API_HOST
    ? import.meta.env.VITE_API_HOST
    : window.location.host;

export default function useWebSocket(socket, messageFn) {
    return new Promise((resolve, reject) => {
        const ws = new WebSocket(wsProtocols + "//" + host + "/ws");

        ws.onopen = () => {
            socket.value = ws;
            console.log("ws connected");
            resolve();
        };

        ws.onerror = (err) => {
            socket.value = null;
            console.log("ws error", err);
            reject(err);
        };

        ws.onclose = () => {
            socket.value = null;
            console.log("ws disconnected");
            reject();
        };

        ws.onmessage = (event) => {
            let data = JSON.parse(event.data);
            messageFn(data);
        };
    });
}
