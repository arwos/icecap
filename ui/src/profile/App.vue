<script setup lang="ts">
import UserForm from '@/profile/UserForm.vue'
</script>

<template>
  <Suspense>
    <template #default>
      <aside class="bq bq-primary area">
        <UserForm />
      </aside>
    </template>

    <template #fallback>
      <aside class="bq bq-warning area">
        <div>Loading...</div>
      </aside>
    </template>
  </Suspense>
</template>

<style scoped lang="scss"></style>
