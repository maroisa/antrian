import { createSignal } from "solid-js";

const [loketData, setLoketData] = createSignal({
    A: 0,
    B: 0,
    C: 0,
    D: 0,
    E: 0,
    F: 0,
});

const [newLoket, setNewLoket] = createSignal(null);

const wsProtocols = window.location.protocol == "http:" ? "ws:" : "wss:";
const ws = new WebSocket(
    wsProtocols + "//" + import.meta.env.VITE_API_HOST + "/loket/ws",
);

ws.onopen = () => console.log("ws connected");
ws.onclose = () => {
    console.log("ws disconnected");
    setTimeout(() => window.location.reload(), 1000);
};

ws.onmessage = (event) => {
    const data = JSON.parse(event.data);
    const changedKey = Object.keys(data).find(
        (key) => loketData()[key] !== data[key],
    );
    const result = changedKey ? changedKey + data[changedKey] : "C0";
    setNewLoket(result);

    setLoketData(data);
};

export { ws as socket, loketData, newLoket };
