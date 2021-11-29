<template>
  <div>
   <v-row no-gutters>
     <v-col sm="3" class="pt-13 justify-center align-center">
       <h1 style="font-weight: bolder;">Bestellübersicht</h1>
     </v-col>
     <v-spacer></v-spacer>
     <v-col sm="2" class="pt-15 justify-center align-center">
       <v-select
           :items="items"
           filled
           outlined
           label="Tage"
           solo
       ></v-select>
     </v-col>
   </v-row>
   <v-row no-gutters>
     <v-toolbar>
       <v-tabs
           dark
           background-color="grey"
           outlined
           grow
       >
         <v-tab>
           <v-badge
               color="green"
               outlined
               content="35">
             Bestellungen
           </v-badge>
         </v-tab>

         <v-tab>
           <v-badge
               color="green"
               content="94%"
               outlined
           >
             bestellungen Bestatigt
           </v-badge>
         </v-tab>

         <v-tab>
           <v-badge
               color="green"
               content="5"
               outlined
           >
             bestellungen/Uhr
           </v-badge>
         </v-tab>
       </v-tabs>
     </v-toolbar>
     <v-col cols="12">
       <Bestellübersicht class="mt-10"></Bestellübersicht>
     </v-col>
     <v-container fluid>
       <v-row align="center">
         <v-col cols="6">
           <v-subheader>
             Bestellungen
           </v-subheader>
         </v-col>

         <v-col cols="6">
           <v-select
               v-model="select"
               :items="amount"
               item-text="state"
               label="Beim Abschnitt"
               filled
               outlined
               persistent-hint
               return-object
               single-line
           ></v-select>
         </v-col>
       </v-row>
     </v-container>
     <v-sparkline
         :value="value"
         :gradient="gradient"
         :smooth="radius || false"
         :padding="padding"
         :line-width="width"
         :stroke-linecap="lineCap"
         :gradient-direction="gradientDirection"
         :fill="fill"
         :type="type"
         :auto-line-width="autoLineWidth"
         auto-draw
     ></v-sparkline>
   </v-row>
    <v-card-actions>
      <v-row>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              block
              rounded
              color="red"
              dark
          >
            Pdf Export
          </v-btn>
        </v-col>
      </v-row>

    </v-card-actions>
   </div>
</template>

<script>
const gradients = [
 ['#222'],
 ['#42b3f4'],
 ['red', 'orange', 'yellow'],
 ['purple', 'violet'],
 ['#00c6ff', '#F0F', '#FF0'],
 ['#f72047', '#ffd200', '#1feaea'],
]
export default {
 data: () => ({


     amount: [
       { state: 'Sample1',},
       { state: 'Sample2',},
       { state: 'sample3',},
       { state: 'sample4',},
       { state: 'sample5',},
     ],

   items: ['Heute', 'Gestern','Vor Gestern',],

   width: 2,
   radius: 10,
   padding: 8,
   lineCap: 'round',
   gradient: gradients[5],
   value: [0, 5, 8, 20, 5, 10, 3, 5, 1, 8, 2, 9, 5, 0,],
   gradientDirection: 'top',
   gradients,
   fill: false,
   type: 'trend',
   autoLineWidth: false,
 }),
}
</script>
