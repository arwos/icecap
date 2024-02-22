<script setup lang="ts">
import type { ISPModel } from '@/components/oauth/types'

defineProps<ISPModel>()
</script>

<template>
  <a class="btn btn-secondary" :href="url">
    <img :src="icon" alt="" class="logo">
    <span class="ml-1">{{ title }}</span>
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