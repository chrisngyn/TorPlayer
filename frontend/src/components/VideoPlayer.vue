<script setup lang="ts">
import { ref, watch } from "vue";
import { arrayBufferToArrayNumber, b64toBlob, getFileExtension } from "@/ultis";
import { StandardizeSubtitle } from "../../wailsjs/go/main/App";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

const props = defineProps<{
  title: string;
  src: string;
  subtitles?: { label: string; url: string; }[];
}>();

const videoRef = ref<HTMLVideoElement | null>(null);
const subtitles = ref<{ label: string; url: string; }[]>(props.subtitles ?? []);
const selectedSubtitle = ref<{
  label: string;
  source: string | File;
} | null>(null);


const subSyncMilliseconds = ref<number>(0);


watch(selectedSubtitle, async () => {
  // reset subSyncMilliseconds
  subSyncMilliseconds.value = 0;
  // sync subtitle
  await updateSubtitle();
});

const addSyncAmount = async (amount: number) => {
  if (!selectedSubtitle.value) return;
  subSyncMilliseconds.value += amount;

  await updateSubtitle();
};

const resetSyncAmount = async () => {
  subSyncMilliseconds.value = 0;

  await updateSubtitle();
};

async function updateSubtitle() {
  removeAllSubtitles();
  await syncSubtitle();
}

let textTrackObjectURL = "";

function removeAllSubtitles() {
  for (const track of videoRef.value?.querySelectorAll("track") ?? []) videoRef.value?.removeChild(track);

  if (textTrackObjectURL) URL.revokeObjectURL(textTrackObjectURL);
}

async function syncSubtitle() {
  if (!selectedSubtitle.value) return;

  const sub = selectedSubtitle.value;

  console.log(`Syncing subtitle to ${sub.label} from ${typeof sub.source === "string" ? sub?.source : sub?.source.name}`);

  const ext = getFileExtension(typeof sub.source === "string" ? sub?.source : sub?.source.name);
  let content: ArrayBuffer;
  if (typeof sub.source === "string") {
    const resp = await fetch(sub.source);
    content = await resp.arrayBuffer();
  } else {
    content = await sub.source.arrayBuffer();
  }
  // go function is []byte, but return here is base64 string, so cast to any.
  const subContent = await StandardizeSubtitle(
    arrayBufferToArrayNumber(content),
    ext,
    subSyncMilliseconds.value || 0,
  ) as any;

  const textTrackUrl = URL.createObjectURL(b64toBlob(subContent, "text/vtt"));

  console.log(textTrackUrl);

  const textTrack = document.createElement("track");
  textTrack.kind = "subtitles";
  textTrack.label = sub.label;
  textTrack.srclang = "vn";
  textTrack.src = textTrackUrl;
  textTrack.default = true;

  // update textTrackObjectURL
  textTrackObjectURL = textTrackUrl;

  videoRef.value?.appendChild(textTrack);
}

interface HTMLInputEvent extends Event {
  target: HTMLInputElement & EventTarget;
}

const onFileInputChange = async (ev: Event) => {
  let files = (ev as HTMLInputEvent).target.files || (ev as DragEvent).dataTransfer?.files;
  if (!files?.length) {
    return;
  }

  const file = files[0];

  selectedSubtitle.value = {
    label: `User uploaded: ${file.name}`,
    source: file,
  };
};


</script>

<template>
  <section>
    <video
      ref="videoRef"
      controls autoplay disablepictureinpicture
      preload="auto" crossorigin="anonymous"
      class="w-full my-2"
      :title="$props.title">
      <source :src="$props.src" />
    </video>
    <div class="subtitle-controller mt-4 flex flex-col">
      <div class="flex justify-between items-center">
        <h3 class="text-xl font-bold my-4"><span class="border-b-2 border-red-700 pb-1">Subtitles</span></h3>
        <p class="text-sm pr-1">{{ selectedSubtitle?.label }}</p>
      </div>

      <div class="grid grid-cols-2 gap-1">
        <div class="col-span-1 flex flex-col py-3 px-1 rounded-sm hover:bg-stone-700">
          <h4 class="mb-1 font-thin text-sm">Adjustment</h4>
          <div class="flex items-center">
            <button class="px-4 py-2 m-1 rounded bg-stone-900 hover:bg-red-700" @click="resetSyncAmount()">
              <font-awesome-icon icon="fa-solid fa-eraser" /> Reset
            </button>
            <button class="px-4 py-2 m-1 rounded bg-stone-900 hover:bg-red-700" @click="addSyncAmount(-500)">
              <font-awesome-icon icon="fa-solid fa-backward" /> -0.5s
            </button>
            <p class="mx-2">{{ subSyncMilliseconds / 1000 }}s</p>
            <button class="px-4 py-2 ml-2 rounded bg-stone-900 hover:bg-red-700" @click="addSyncAmount(500)">
              <font-awesome-icon icon="fa-solid fa-forward" /> +0.5s
            </button>
          </div>
        </div>
        <div class="col-span-1 flex flex-col py-3 px-1 rounded-sm hover:bg-stone-700">
          <h4 class="mb-1 font-thin text-sm">Add sub from file</h4>
          <input
            type="file"
            name="fileInput"
            accept=".vtt, .srt, application/x-subrip, text/vtt"
            @input="onFileInputChange"
            class="border text-stone-100 border-gray-400 hover:border-red-700 file:mr-5 rounded bg-stone-700 file:border-[0px] file:p-4 file:bg-stone-700 file:rounded-l file:text-stone-100 hover:file:cursor-pointer hover:file:bg-stone-800 hover:file:text-red-700"
          />
        </div>
      </div>
      <div class="mt-2 rounded-sm">
        <h4 class="px-1 py-1 mb-1 text-lg">Available subtitles</h4>
        <div class="flex  gap-2 flex-wrap">
          <template v-for="sub in subtitles" :key="sub.url">
            <template v-if="sub.url == selectedSubtitle?.source">
              <button
                class="px-4 py-2 rounded bg-red-600 hover:bg-red-700 text-slate-100"
                @click="selectedSubtitle = null"
              >
                {{ sub.label }}
              </button>
            </template>
            <template v-else>
              <button
                class="px-4 py-2 rounded bg-stone-900 hover:bg-stone-800 text-slate-100 "
                @click="selectedSubtitle = { label: sub.label, source: sub.url }"
              >
                {{ sub.label }}
              </button>
            </template>
          </template>
        </div>
      </div>
    </div>
  </section>
</template>