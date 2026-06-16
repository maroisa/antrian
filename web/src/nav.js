import { SpeakerWaveIcon } from "@heroicons/vue/24/outline";
import { ComputerDesktopIcon } from "@heroicons/vue/24/outline";
import { PlusCircleIcon } from "@heroicons/vue/24/outline";

const NAV = [
    {
        to: "/panggil-tanpa-daftar",
        name: "panggil Tanpa Daftar",
        color: "secondary",
        icon: SpeakerWaveIcon,
    },
    {
        to: "/panggil",
        name: "Panggil",
        color: "primary",
        icon: SpeakerWaveIcon,
    },
    {
        to: "/daftar",
        name: "Daftar",
        color: "accent",
        icon: PlusCircleIcon,
    },
];
export default NAV;
