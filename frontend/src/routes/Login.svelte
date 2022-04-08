<script>
    import cymruDragonLogo from "../assets/images/wails-logo-horizontal-dark.png";
    let submit = false;
    let invalidPassword = false;
    let unlockingKeys = false;
    let passwordLabel = "password"
    let password = ""

    const handleSubmit = () => {
        
        console.log("Submitted...")
        submit = true;
        unlockingKeys = true;
        sleep(2000).then(() => {
            console.log("Cancelling unlock...")
        unlockingKeys = false;
        });
        // do logic, submit val is irrelicent.
        doTheGreetThing()
    }
    function sleep(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }

    function doTheGreetThing() {
        // @ts-ignore
        window.go.raverte.Raverte.GetChartOnlyMode().then((/** @type {bool}  */ result) => {
            console.log(result)
        });
    }

    /*
indow.go.main.App.Greet(name).then((result) => {
    greeting = result;
		});
    logUserin() {
                this.unlockingKeys = true;
                window.backend.Raverte.UnlockKeys(this.$refs.pword.value).then(() => {
                    this.$router.replace("/mainstage")
                }).catch((error) => {
                    if (error == "invalid password") {
                        // TODO: if x failed, offer to delete.
                        this.passwordFieldMsg = error
                        this.unlock = "Unlock";
                        this.unlockingKeys = false;
                        this.invalidPassword = true;
                    }else{
                        this.passwordFieldMsg = "password" // incase we entered incorrect password first, reset
                        this.invalidPassword = false;
                        console.log(error)
                        // TODO: error mechanism.
                    }
                });
            } */
</script>

<svelte:head>
    <title>
        Raverte - Login
    </title>
</svelte:head>

<div class="h-screen bg-canvas flex flex-col items-center justify-center mx-auto">
    <div class="shadow-lgreen min-w-[420px] rounded-[15px]">
        <div class="h-[500px] shadow-lred rounded-[15px]">
            <div class="h-[440px] grid grid-cols-7 bg-canvas rounded-[15px]" data-wails-no-drag>
                <p class="my-14 col-start-1 col-end-8 text-white text-center text-4xl font-poppins font-light">RAVERTE</p>
                <form class="col-start-2 col-end-7" autocomplete="off" on:submit|preventDefault={handleSubmit}>
                    <div class="flex flex-col-reverse pt-1">
                        <input type="password" class="key-password-input text-white" required />
                        <label class="block font-poppins font-light text-white {invalidPassword === true ? 'invalid-password' : 'password-label'}" for="password">
                            {passwordLabel} <!-- focus:translate-y-12 -->
                        </label>
                    </div>
                    <div class="flex justify-end pt-1 cursor-pointer">
                        <p class="w-5/12 text-white font-poppins font-extralight text-right text-xs">
                            Forgot Password?
                        </p>
                    </div>
                    <div class="flex justify-center pt-20">
                        <div class="{unlockingKeys === true ? 'unlocking': 'hidden'}"></div>
                        <button type="submit" class="{unlockingKeys === true ? 'hidden' : 'w-6/12 h-11 p-1 leading-10 rounded-[30px] block font-poppins font-normal transition ease-in-out duration-300 text-lbtndark hover:text-white bg-gradient-to-tr from-rred to-rgreen hover:shadow-lbtn'}">UNLOCK</button>
                    </div>
                </form>              
            </div>

            <div class="h-[60px]">
                <p class="text-white font-poppins font-extralight text-center text-xs">
                    Powered By
             </p>
            <img src="{cymruDragonLogo}" class="w-32 mx-auto" alt="WAILS">
            </div>   
        </div>
    </div>
</div>
<!--
     <form id="login-form" @submit.prevent="submit" autocomplete="off">
            <div class="row mt-2">
                <div class="col-lg-10 offset-lg-1 mt-5 pword">
                    <input type="password" ref="pword" required />
                    <label v-bind:class="[invalidPassword ? 'invalid-password' : 'password-label']" for="password"> {{ passwordFieldMsg }}</label>
                </div>
            </div>
            <div class="row">
                <div class="col-lg-4 offset-lg-7 mt-2 forgot-password">

                    Forgot Password?
                </div>
            </div>
            <div class="row mt-3">
                <div class="mt-5">
                    <div v-bind:class="[unlockingKeys ? 'unlocking' : 'hidden']"> </div>
                    <button type="button" v-bind:class="[unlockingKeys ? 'hidden' : 'unlock']" v-on:click="logUserin()">UNLOCK</button>        
                </div>
            </div>
        </form>
-->