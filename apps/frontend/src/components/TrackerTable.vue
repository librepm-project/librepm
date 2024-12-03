<template>
    <v-table>
      <thead>
        <tr>
          <th class="text-left">
            Name
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in items"
          :key="item.id"
        >
          <td>{{ item.name }}</td>
        </tr>
      </tbody>
    </v-table>
</template>

<script setup lang="ts">
import { Tracker } from '@/lib/interfaces/tracker.interface';


defineProps<{
    items: Tracker[]
}>()

</script>