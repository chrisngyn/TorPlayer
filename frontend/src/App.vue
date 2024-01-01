<script setup lang="ts"></script>

<template>
  <div class="flex flex-col min-h-screen w-full box-border bg-stone-800">
    <nav class="flex justify-between items-center px-4 py-1 bg-stone-900 text-slate-100">
      <ul class="flex">
        <li class="mr-6 p-2">
          <router-link class="hover:text-red-700" to="/">Home</router-link>
        </li>
        <li class="mr-6 p-2">
          <router-link class="hover:text-red-700" to="/about">About</router-link>
        </li>
<!--                <li class="mr-6 p-2">-->
<!--                  <router-link class="hover:text-red-700" to="/t/123/files">Video</router-link>-->
<!--                </li>-->
      </ul>
    </nav>
    <div class="flex-grow min-w-full text-slate-100">
      <RouterView />
    </div>
  </div>
</template>

<style scoped>
li:has(> a.router-link-active) {
  @apply border-b-2 border-red-700
}
</style>
