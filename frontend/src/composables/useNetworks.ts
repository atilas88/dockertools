import {mockResponse} from '../mocks/delay'
import {networks as mockNetworks} from '../mocks/resources'
import type {Network} from '../types/docker'
import {createListStore} from './createListStore'

// TODO etapa 8: return ListNetworks() from wailsjs/go/main/App.
export const useNetworks = createListStore<Network>(() => mockResponse(mockNetworks))
