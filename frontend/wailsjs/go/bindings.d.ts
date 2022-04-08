import * as models from './models';

export interface go {
  "raverte": {
    "Raverte": {
		GetChartOnlyMode():Promise<boolean>
		LoadProfile():Promise<void>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
