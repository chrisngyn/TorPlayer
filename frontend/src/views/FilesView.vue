<script setup lang="ts">

import { useRoute, useRouter } from "vue-router";
import { onMounted, ref } from "vue";
import { torrent } from "../../wailsjs/go/models";
import { GetTorrentInfo } from "../../wailsjs/go/main/App";
import { bytesLengthToSize, isVideoFile } from "@/ultis";
import BackButton from "@/components/BackButton.vue";

const route = useRoute();
const router = useRouter();
const infoHash = typeof route.params.infoHash == "string" ? route.params.infoHash : route.params.infoHash[0];

const loading = ref(true);
const torrentInfo = ref<torrent.Info | null>(null);

onMounted(async () => {
  console.log(route.params.infoHash);
  try {
    torrentInfo.value = await GetTorrentInfo(infoHash);
  } finally {
    loading.value = false;
  }
});

const watchVideo = async (fileName: string) => {
  await router.push({ name: "watch", params: { infoHash: route.params.infoHash, fileName: fileName } });
};

</script>

<template>
  <main class="w-full min-h-full p-4">
    <div class="flex justify-between">
      <p class="text-sm italic">Info Hash: {{ $route.params.infoHash }}</p>
      <BackButton />
    </div>
    <div v-if="loading">Loading...</div>
    <template v-else>
      <template v-if="torrentInfo">
        <h3 class="text-4xl my-4">{{ torrentInfo.name }}</h3>
        <p>Size: {{ bytesLengthToSize(torrentInfo.length) }}</p>
        <table class="table-auto w-full my-4 border-2 border-gray-400">
          <thead>
          <tr class="border-gray-400 border-b-2">
            <th class="px-4 py-2">Name</th>
            <th class="px-4 py-2">Length</th>
            <th class="px-4 py-2">Completed</th>
            <th class="px-4 py-2">Action</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="file in torrentInfo.files" :key="file.displayPath" class="hover:text-red-600">
            <td class="px-4 py-2">{{ file.displayPath }}</td>
            <td class="px-4 py-2 text-right">{{ bytesLengthToSize(file.length) }}</td>
            <td class="px-4 py-2 text-right">{{ (file.bytesCompleted / file.length).toFixed(2) * 100 }}%</td>
            <td class="px-4 py-2 text-center">
              <button v-if="isVideoFile(file.displayPath)" class="px-4 py-2 bg-red-600 rounded-sm text-slate-100"
                      @click="() => watchVideo(file.displayPath)">
                <font-awesome-icon icon="fa-play" />
                Watch
              </button>
            </td>
          </tr>
          </tbody>
        </table>
      </template>
    </template>
  </main>
</template>
