export namespace docker {
	
	export enum ContainerState {
	    CREATED = "created",
	    DEAD = "dead",
	    EXITED = "exited",
	    PAUSED = "paused",
	    REMOVING = "removing",
	    RESTARTING = "restarting",
	    RUNNING = "running",
	}
	export class Container {
	    id: string;
	    name: string;
	    image: string;
	    state: ContainerState;
	    exitCode?: number;
	    ports: string[];
	    created: number;
	
	    static createFrom(source: any = {}) {
	        return new Container(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.image = source["image"];
	        this.state = source["state"];
	        this.exitCode = source["exitCode"];
	        this.ports = source["ports"];
	        this.created = source["created"];
	    }
	}
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

