export namespace model {
	
	export class UsuarioCadastrado {
	    nome: string;
	    email: string;
	    senha: string;
	
	    static createFrom(source: any = {}) {
	        return new UsuarioCadastrado(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.nome = source["nome"];
	        this.email = source["email"];
	        this.senha = source["senha"];
	    }
	}

}

