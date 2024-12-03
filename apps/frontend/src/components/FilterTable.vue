<template>
    <v-table>
      <thead>
        <tr>
          <th class="text-left">
            Name
          </th>
          <th class="text-left">
            Description
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in items"
          :key="item.name"
        >
          <td>{{ item.name }}</td>
          <td>{{ item.description }}</td>
        </tr>
      </tbody>
    </v-table>
</template>

<script setup lang="ts">
import { Filter } from '@/lib/interfaces/filter.interface';


defineProps<{
    items: Filter[]
}>()

</script>