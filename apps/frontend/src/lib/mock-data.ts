import { Board } from './interfaces/board.interface';
import { Filter } from './interfaces/filter.interface';
import { Issue } from './interfaces/issue.interface';
import { User } from './interfaces/user.interface';

export const boardMock: Board[] = [
  {
    id: 'abc',
    name: 'Board 1',
    description: 'Lorem ipsum dolor amet',
  },
  {
    id: 'cba',
    name: 'Board 2',
    description: 'Lorem ipsum dolor amet',
  },
];

export const filterMock: Filter[] = [
  {
    id: 'abc',
    name: 'Filter 1',
    conditions: [],
    description: 'Lorem ipsum dolor amet',
  },
  {
    id: 'cba',
    name: 'Filter 2',
    conditions: [],
    description: 'Lorem ipsum dolor amet',
  },
];

export const issueMock: Issue[] = [
  {
    key: 'SW-1',
    summary: 'First Issue',
    status: 'In progress',
    description: 'Lorem ipsum dolor amet',
  },
  {
    key: 'SW-2',
    summary: 'Second Issue',
    status: 'To Do',
    description: 'Lorem ipsum dolor amet',
  },
];

export const userMock: User[] = [
  {
    id: 'abc',
    email: 'example1@company.org',
    firstName: 'John',
    lastName: 'Doe',
  },
  {
    id: 'cba',
    email: 'example2@company.org',
    firstName: 'Joe',
    lastName: 'Smish',
  },
];
