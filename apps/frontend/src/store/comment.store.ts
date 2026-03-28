import { defineStore } from 'pinia';
import { Comment } from '@/lib/interfaces/comment.interface';
import commentApi from '@/api/comment.api';

interface CommentStore {
  comments: Comment[];
}

export const useCommentStore = defineStore('comment', {
  state: (): CommentStore => ({
    comments: [],
  }),
  actions: {
    async get(issueId: string) {
      this.comments = await commentApi.indexByIssue(issueId);
    },
    async create(issueId: string, content: string) {
      const created = await commentApi.create(issueId, content);
      this.comments.push(created);
      return created;
    },
    async update(commentId: string, content: string) {
      const updated = await commentApi.update(commentId, content);
      const idx = this.comments.findIndex(c => c.id === commentId);
      if (idx !== -1) this.comments[idx] = updated;
      return updated;
    },
    async destroy(commentId: string) {
      await commentApi.destroy(commentId);
      this.comments = this.comments.filter(c => c.id !== commentId);
    },
  },
});
