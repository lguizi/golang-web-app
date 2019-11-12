function voltar() {
    opcoes.innerHTML = `<p>MENU:</p>
    <p><input type="button" value="1. LISTAR CONTATOS" id="listar" onclick="lista()"></p>
    <p><input type="button" value="2. CRIAR CONTATOS" id="criar" onclick="criar()"><p>
    <p><input type="button" value="3. EDITAR CONTATOS" id="editar" onclick="editar()"></p>
    <p><input type="button" value="4. DELETAR CONTATOS" id="del" onclick="del()"></p>`
}

function lista() {
    
    opcoes.innerHTML = `<p>LISTAR CONTATOS:</p>
    <p><select name="sel" id="sel" size="10"></select></p>
    <p><input type="button" value="Voltar" id="voltar" onclick='voltar()'></p>`
}

function criar() {
    opcoes.innerHTML = `<p>CRIAR CONTATOS:</p>
    <p>Nome: <input type="text" name="nome" id="nome" maxlength="45" required></p>
    <p>Celular: <input type="tel" name="tel" id="tel" maxlength="11" required></p>
    <p>E-mail: <input type="email" name="email" id="email" maxlength="45" required></p>
    <p><input type="button" value="Cadastrar" id="cadastrar" onclick='cadastra()'>
    <input type="button" value="Voltar" id="voltar" onclick='voltar()'></p>`
}

function editar() {
    opcoes.innerHTML = `<p>EDITAR CONTATOS:</p>
    <p><select name="sel" id="sel" size="10"></select></p>
    <p><input type="button" value="Editar" id="editar" onclick='editar()'>
    <input type="button" value="Voltar" id="voltar" onclick='voltar()'></p>`
}

function del() {
    opcoes.innerHTML = `<p>DELETAR CONTATOS:</p>
    <p><select name="sel" id="sel" size="10"></select></p>
    <p><input type="button" value="Deletar" id="del" onclick='del()'>
    <input type="button" value="Voltar" id="voltar" onclick='voltar()'></p>`
}

let opcoes = document.querySelector(`div`)
