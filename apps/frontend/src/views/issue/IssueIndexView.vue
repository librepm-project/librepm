<template>
  <div>
    <h1>Issues</h1>
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
          v-for="item in issues"
          :key="item.key"
        >
          <td>{{ item.key }}</td>
          <td>{{ item.summary }}</td>
          <td>{{ item.status }}</td>
        </tr>
      </tbody>
    </v-table>
  </div>
</template>

<script lang="ts" setup>
  import Issue from '../../lib/interfaces/issue.interface.ts';
  
  const issues: Issue[] = [
    {
      key: "SW-1",
      summary: "First Issue",
      status: "In progress"
    },
    {
      key: "SW-2",
      summary: "Second Issue",
      status: "To Do"
    }
  ];

</script>
