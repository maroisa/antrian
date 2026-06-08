import { A, useParams } from "@solidjs/router";
import LoketMini from "../components/LoketView";
import { ChevronLeft } from "../components/Svg";
import { socket } from "../utils/socket";

export default function LoketItem() {
    const params = useParams();
    return (
        <main class="flex flex-col-reverse md:flex-row">
            <div class={`grow w-full p-2 border-2 bg-${params.kode}`}>
                <div class="bg-white p-2 shadow-md h-full border-2 ">
                    <A
                        href="/"
                        class={`btn rounded-none border-2 border-black bg-${params.kode}`}
                    >
                        <ChevronLeft className="size-4" />
                        Kembali
                    </A>
                    <p>ini adalah Loket {params.kode}</p>
                    <button
                        onClick={() => {
                            socket.send(params.kode);
                        }}
                        class="btn"
                    >
                        Antrian Selanjutnya
                    </button>
                </div>
            </div>
            <LoketMini />
        </main>
    );
}
