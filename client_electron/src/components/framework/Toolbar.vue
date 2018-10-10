<template>
    <!-- Hero head: will stick at the top -->
    <div class="hero-head" style="border-bottom: 5px solid #e50112;">
        <nav class="navbar ">
            <div class="navbar-brand">
                <a class="navbar-item" @click="showDrawer">
                    <b-icon icon="view-headline"></b-icon>
                </a>
            </div>
            <div class="navbar-menu">
                <div class="navbar-end">
                    <div class="navbar-item" style="font-size: 14px;">
                        <i class="mdi mdi-eye mdi-24px" style="color: rgb(125, 125, 125);"></i>
                    </div>
                </div>
            </div>
        </nav>

    </div>
</template>
 
<script>  
export default {
    data() {
        return {
        }
    },
    methods: {
        showDrawer() {
            if (this.$globalvar.configMode) {
                this.$globalvar.toastInfo("You are in Configuration Mode.");
            } else {
                this.$globalEventHub.$emit("showDrawer");
            }
        }
    }
}
</script>