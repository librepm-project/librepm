<template>
    <v-table>
      <thead>
        <tr>
          <th class="text-left">
            Key
          </th>
          <th class="text-left">
            Summary
          </th>
          <th class="text-left">
            Status
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in items"
          :key="item.key"
        >
          <td>{{ item.key }}</td>
          <td>{{ item.summary }}</td>
          <td>{{ item.status }}</td>
        </tr>
      </tbody>
    </v-table>
</template>

<script setup lang="ts">

defineProps<{
    items: Issue[]
}>()

</script>