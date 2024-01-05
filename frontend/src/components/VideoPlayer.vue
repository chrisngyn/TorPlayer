<script setup lang="ts">
import { ref } from "vue";
import { arrayBufferToArrayNumber, b64toBlob, getFileExtension } from "@/ultis";
import { StandardizeSubtitle } from "../../wailsjs/go/main/App";

const props = defineProps<{
  title: string;
  src: string;
  subtitles?: { label: string; url: string; }[];
}>();

const videoRef = ref<HTMLVideoElement | null>(null);
const subtitles = ref<{ label: string; url: string; }[]>(props.subtitles ?? []);
const selectedSubtitle = ref<{
  label: string;
  originUrl: string;
  objectUrl: string;
} | null>(null);


function removeAddSubtitles() {
  for (const track of videoRef.value?.querySelectorAll("track") ?? []) videoRef.value?.removeChild(track);

  if (!selectedSubtitle.value) return;
  URL.revokeObjectURL(selectedSubtitle.value.objectUrl);
  selectedSubtitle.value = null;
}

async function changeSubtitle(label: string, url: string, srclang?: string) {
  console.log(`Changing subtitle to ${label} at ${url}`);
  removeAddSubtitles();

  const ext = getFileExtension(url);
  const resp = await fetch(url);
  // go function is []byte, but return here is base64 string, so cast to any.
  const subContent = await StandardizeSubtitle(arrayBufferToArrayNumber(await resp.arrayBuffer()), ext, 0) as any;

  const textTrackUrl = URL.createObjectURL(b64toBlob(subContent, "text/vtt"));

  console.log(textTrackUrl);

  const textTrack = document.createElement("track");
  textTrack.kind = "subtitles";
  textTrack.label = label;
  textTrack.srclang = srclang ?? "vn";
  textTrack.src = textTrackUrl;
  textTrack.default = true;
  videoRef.value?.appendChild(textTrack);
  selectedSubtitle.value = { label, originUrl: url, objectUrl: textTrackUrl };
}

</script>

<template>
  <section>
    <video
      ref="videoRef"
      controls autoplay disablepictureinpicture preload crossorigin="anonymous"
      class="w-full my-2"
      :title="$props.title">
      <source :src="$props.src" />
    </video>
    <div class="subtitle-controller mt-4">
      <h3 class="text-lg my-4"><span class="border-b-2 border-red-700 pb-1">Subtitles</span></h3>
      <div class="-mx-1">
        <template v-for="sub in subtitles" :key="sub.url">
          <template v-if="sub.url == selectedSubtitle?.originUrl">
            <button
              class="px-4 py-2 m-1 rounded bg-red-600 hover:bg-red-700 text-slate-100"
              @click="removeAddSubtitles()"
            >
              {{ sub.label }}
            </button>
          </template>
          <template v-else>
            <button
              class="px-4 py-2 m-1 rounded bg-stone-900 hover:bg-stone-800 text-slate-100 "
              @click="changeSubtitle(sub.label, sub.url)"
            >
              {{ sub.label }}
            </button>
          </template>
        </template>
      </div>
    </div>
  </section>
</template>