<script setup lang="ts">
import type { TField } from '@/profile/fields/types'

const props = defineProps<TField>()
const emit = defineEmits(['value'])

let _value = props.value
</script>

<template>
  <label class="fs-sm">{{ props.title }}:</label>
  <select v-model="_value" @change="emit('value', props.field, _value)" class="full-width fs-md">
    <option v-for="item in props.data" :key="item" :value="item">{{ item }}</option>
  </select>
</template>

<style scoped lang="scss"></style>
