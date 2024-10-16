<template>
  <div class="about">
    <h1>This is an about page</h1>
    <v-text-field label="Label"></v-text-field>
    <v-btn prepend-icon="$vuetify">Button</v-btn>
  </div>
</template>

<style>
@media (min-width: 768px) {
  .about {
    max-width: 768px;
    margin-left: auto;
    margin-right: auto;
    padding: 0 1rem;
  }
}
</style>
