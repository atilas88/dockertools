import {mockResponse} from '../mocks/delay'
import {volumes as mockVolumes} from '../mocks/resources'
import type {Volume} from '../types/docker'
import {createListStore} from './createListStore'

// TODO etapa 8: return ListVolumes() from wailsjs/go/main/App.
export const useVolumes = createListStore<Volume>(() => mockResponse(mockVolumes))
