<script setup lang="ts">
import type { TField } from '@/profile/fields/types'

const props = defineProps<TField>()
const emit = defineEmits(['value'])

let _value = props.value
</script>

<template>
  <label class="fs-sm">{{ props.title }}:</label>
  <input
    type="text"
    class="full-width fs-md"
    v-model="_value"
    @input="emit('value', props.field, _value)"
  />
</template>

<style scoped lang="scss"></style>
