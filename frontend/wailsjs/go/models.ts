export namespace docker {
	
	export class DockerInfo {
	    version: string;
	    apiVersion: string;
	    host: string;
	
	    static createFrom(source: any = {}) {
	        return new DockerInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.version = source["version"];
	        this.apiVersion = source["apiVersion"];
	        this.host = source["host"];
	    }
	}

}

