import { handleError } from "vue";

const host = import.meta.env.VITE_API_HOST
    ? import.meta.env.VITE_API_HOST
    : window.location.host;

const URL = window.location.protocol + "//" + host + "/";

export async function get(...segments) {
    const [response, error] = await fetch(URL + segments.join("/"))
        .then((res) => [res, null])
        .catch((err) => [null, err]);

    return [response, error];
}
