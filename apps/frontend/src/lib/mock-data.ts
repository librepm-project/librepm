import { Board } from "@/lib/interfaces/board.interface";
import { Filter } from "@/lib/interfaces/filter.interface";
import { Issue } from "@/lib/interfaces/issue.interface";
import { Project } from "@/lib/interfaces/project.interface";
import { Status } from "@/lib/interfaces/status.interface";
import { Tracker } from "@/lib/interfaces/tracker.interface";
import { User } from "@/lib/interfaces/user.interface";


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
    tracker: 'Task',
  },
  {
    key: 'SW-2',
    summary: 'Second Issue',
    status: 'To Do',
    description: 'Lorem ipsum dolor amet',
    tracker: 'Task',
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


export const projectMock: Project[] = [
  {
    id: "id",
    codeName: "SW",
    name: "Mocked Project",
  },
];

export const statusMock: Status[] = [
  {
    id: 'id',
    name: "Mocked Project",
  },
];

export const trackerMock: Tracker[] = [
  {
    id: 'id',
    name: "Mocked Project",
  },
];
