<template>
    <div>
        <h3>Informations médecin</h3>
            <drizzle-contract-form
                contractName="Covid"
                method="setDoctor"
                :placeholders="['Adresse :', 'Nom :', 'IDNumber :']"
            />
        <h3>Informations patient</h3>
            <drizzle-contract-form
                contractName="Covid"
                method="setPatient"
                :placeholders="['Adresse patient :', 'Nom :', 'Email :', 'Age :', 'Genre (1 : masculin, 2 : féminin, 3 : autre) :']"
            />
        <h3>Informations traitement</h3>
            <drizzle-contract-form
                contractName="Covid"
                method="setPatientTreatment"
                :placeholders="['Adresse patient :', 'Description :', 'Médicament :']"
            />
    </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: 'Covid',
  computed: mapGetters('drizzle', ['isDrizzleInitialized'])
}
</script>

<style></style>
