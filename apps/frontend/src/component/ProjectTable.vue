<template>
    <v-table>
      <thead>
        <tr>
          <th class="text-left">
            Code
          </th>
          <th class="text-left">
            Name
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in items"
          :key="item.codeName"
        >
          <td>{{ item.codeName }}</td>
          <td>{{ item.name }}</td>
        </tr>
      </tbody>
    </v-table>
</template>

<script setup lang="ts">
import { Project } from '@/lib/interfaces/project.interface';



defineProps<{
    items: Project[]
}>()

</script>