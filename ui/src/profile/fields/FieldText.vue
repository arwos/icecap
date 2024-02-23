<script setup lang="ts">
import type { TField } from '@/profile/fields/types'
const props = defineProps<TField>()
</script>

<template>
  <label class="fs-sm">{{ props.title }}:</label>
  <p class="field-text">{{ props.value }}</p>
</template>

<style scoped lang="scss">
.field-text {
  border: 1px solid var(--ong-color-secondary-bg);
  padding: var(--ong-padding-y-size) var(--ong-padding-x-size);
}
</style>
