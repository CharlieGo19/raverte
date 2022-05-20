<script>
    //import cymruDragonLogo from "../assets/images/wails-logo-horizontal-dark.png";
    let submit = false;
    let invalidPassword = false;
    let unlockingKeys = false;
    let passwordLabel = "password";
    let password = "";

    
    async function login() {
        unlockingKeys = true;
        // @ts-ignore
        await window.go.raverte.Raverte.UnlockKeys(password).then(() => {
            // We're using go to maintain state.
            unlockingKeys = false;
        }).catch((/** @type string */ error) => {
            
            passwordLabel = error;
            invalidPassword = true;
            unlockingKeys = false;
        })
    }

</script>

<svelte:head>
    <title>
        Raverte - Login
    </title>
</svelte:head>

<!-- TODO: Clean this design! It looks like hot 💩 -->

<div class="h-screen bg-canvas flex flex-col items-center justify-center mx-auto z-40 fixed w-full">
    <div class="shadow-lgreen min-w-[420px] rounded-[15px] z-40">
        <div class="h-[500px] shadow-lred rounded-[15px]">
            <div class="h-[440px] grid grid-cols-7 bg-canvas rounded-[15px]" data-wails-no-drag>
                <p class="my-14 col-start-1 col-end-8 text-white text-center text-4xl font-poppins font-light">RAVERTE</p>
                <form class="col-start-2 col-end-7" autocomplete="off" on:submit|preventDefault={login}>
                    <div class="flex flex-col-reverse pt-1">
                        <input type="password" class="key-password-input text-white" bind:value={password} required />
                        <label class="block font-poppins font-light text-white {invalidPassword === true ? 'invalid-password' : 'password-label'}" for="password">
                            {passwordLabel}
                        </label>
                    </div>
                    <div class="flex justify-end pt-1 cursor-pointer">
                        <p class="w-5/12 text-white font-poppins font-extralight text-right text-xs">
                            Forgot Password? <!-- TODO: Add destroy Keyring feature. -->
                        </p>
                    </div>
                    <div class="flex justify-center pt-20">
                        <div class="{unlockingKeys === true ? 'unlocking': 'hidden'}"></div>
                        <button type="submit" class="{unlockingKeys === true ? 'hidden' : 'w-6/12 h-11 p-1 leading-10 rounded-[30px] block font-poppins font-normal transition ease-in-out duration-300 text-lbtndark hover:text-white bg-gradient-to-tr from-rred to-rgreen hover:shadow-lbtn'}">UNLOCK</button>
                    </div>
                </form>              
            </div>
           
            <!-- Re add this when we have nice design. -->
            <!-- <div class="h-[60px]">
                <p class="text-white font-poppins font-extralight text-center text-xs">
                    Powered By
                </p>
                <img src="{cymruDragonLogo}" class="w-32 mx-auto" alt="WAILS">
            </div> -->
        </div>
    </div>
</div>