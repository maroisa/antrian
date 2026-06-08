import { For } from "solid-js";
import { loketData } from "../utils/socket";
import { A } from "@solidjs/router";
import { Monitor } from "../components/Svg";

export default function HomeView() {
    return (
        <div class="h-full w-full flex justify-center items-center p-4">
            <div class="flex flex-col gap-8 p-8 bg-base-100 rounded-lg shadow-lg w-full max-w-lg">
                <a
                    href="/loket"
                    class="btn btn-soft text-xl font-bold py-8 w-full flex gap-2"
                >
                    <Monitor className="size-6 md:size-8" />
                    Tampilan Penuh
                </a>
                <nav class="w-full grid grid-cols-2 gap-2 *:btn *:min-h-24 *:font-bold *:text-xl *:md:text-2xl">
                    <For each={Object.keys(loketData())}>
                        {(item) => (
                            <A
                                class={`btn shadow-lg rounded-lg bg-${item}`}
                                href={`/loket/${item}`}
                            >
                                Loket {item}
                            </A>
                        )}
                    </For>
                </nav>
            </div>
        </div>
    );
}
