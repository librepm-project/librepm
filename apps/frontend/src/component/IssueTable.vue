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
          <td><router-link :to="`/issue/${item.id}`">{{ item.summary }}</router-link></td>
          <td>{{ item.status.name }}</td>
        </tr>
      </tbody>
    </v-table>
</template>

<script setup lang="ts">
import { Issue } from '@/lib/interfaces/issue.interface';


defineProps<{
    items: Issue[]
}>()

</script>