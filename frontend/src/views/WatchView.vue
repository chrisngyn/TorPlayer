<script setup lang="ts">
import { useRoute } from "vue-router";
import { onMounted, onUnmounted, ref } from "vue";
import { GetTorrentInfo, StartDownload, StopDownload } from "../../wailsjs/go/main/App";
import BackButton from "@/components/BackButton.vue";
import VideoPlayer from "@/components/VideoPlayer.vue";

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
  const subtitleFiles = info.files.filter((file) =>
    file.displayPath.endsWith(".srt") ||
    file.displayPath.endsWith(".vtt") ||
    file.displayPath.endsWith(".ass") ||
    file.displayPath.endsWith(".ssa")
  );
  for (const file of subtitleFiles) subtitles.value.push({
    label: file.displayPath,
    url: encodeURI(`/stream/${infoHash}/file/${file.displayPath}`),
  });
});

onUnmounted(async () => {
  console.log(`Stopping download infoHash: ${infoHash}, fileName: ${fileName}`);
  await StopDownload(infoHash, fileName);
});

</script>

<template>
  <main class="w-full p-4 box-border">
    <div class="flex justify-between my-1">
      <h1 class="text-2xl ">{{ fileName }}</h1>
      <BackButton />
    </div>
    <VideoPlayer class="mt-1" :title="fileName" :src="videoSrc" :subtitles="subtitles" />
  </main>
</template>