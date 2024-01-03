<script setup lang="ts">
import { useRoute } from "vue-router";
import { onMounted, onUnmounted, ref } from "vue";
import { GetTorrentInfo, StartDownload, StopDownload } from "../../wailsjs/go/main/App";
import BackButton from "@/components/BackButton.vue";
import toWebVTT from "srt-webvtt";

const route = useRoute();
const infoHash = typeof route.params.infoHash == "string" ? route.params.infoHash : route.params.infoHash[0];
const fileName = typeof route.params.fileName == "string" ? route.params.fileName : route.params.fileName[0];

const videoSrc = encodeURI(`/stream/${infoHash}/file/${fileName}`);
const subtitles = ref<{ label: string; url: string; }[]>([]);

onMounted(async () => {
  console.log(`Starting download infoHash: ${infoHash}, fileName: ${fileName}`);
  await StartDownload(infoHash, fileName);

  console.log("Starting download subtitles");
  const info = await GetTorrentInfo(infoHash);
  const subtitleFiles = info.files.filter((file) => file.displayPath.endsWith(".srt"));
  for (const file of subtitleFiles) subtitles.value.push({
    label: file.displayPath,
    url: encodeURI(`/stream/${infoHash}/file/${file.displayPath}`),
  });
});

onUnmounted(async () => {
  console.log(`Stopping download infoHash: ${infoHash}, fileName: ${fileName}`);
  await StopDownload(infoHash, fileName);
});

async function changeSubtitle(label: string, url: string) {
  console.log(`Changing subtitle to ${label}, ${url}`);
  removeSubtitle();

  const resp = await fetch(url);
  const textTrackUrl = await toWebVTT(await resp.blob());

  const textTrack = document.createElement("track");
  textTrack.kind = "subtitles";
  textTrack.label = label;
  textTrack.srclang = "en";
  textTrack.src = textTrackUrl;
  textTrack.default = true;
  const video = document.querySelector("video");
  video?.appendChild(textTrack);
}

function removeSubtitle() {
  const video = document.querySelector("video");
  video?.querySelectorAll("track").forEach((track) => video.removeChild(track));
}


</script>

<template>
  <main class="w-full">
    <div class="flex justify-between m-4">
      <h1 class="text-xl ">{{ fileName }}</h1>
      <BackButton />
    </div>
    <video controls class="w-full">
      <source
        :src="videoSrc"
        type="video/mp4"
      />
    </video>
    <div>
      <template v-for="sub in subtitles" :key="sub.url">
        <button
          class="bg-stone-900 hover:bg-stone-800 text-slate-100 px-4 py-2 m-2 rounded"
          @click="changeSubtitle(sub.label, sub.url)"
        >
          {{ sub.label }}
        </button>
      </template>
    </div>
  </main>
</template>