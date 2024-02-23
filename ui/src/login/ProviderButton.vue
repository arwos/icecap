<script setup lang="ts">
import type { TProvider } from '@/login/types'

const props = defineProps<TProvider>()
</script>

<template>
  <a class="btn btn-light" :href="props.url">
    <img :src="props.icon" alt="" class="logo" />
    <span class="ml-1">{{ props.title }}</span>
  </a>
</template>

<style scoped lang="scss">
.btn {
  min-height: 24px;
  padding-right: 24px !important;
  padding-left: 24px !important;
}

.logo {
  height: 24px;
  border-radius: 24px;
}
</style>
