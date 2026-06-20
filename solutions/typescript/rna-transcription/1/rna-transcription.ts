export function toRna(rna: string): string {

  if (rna.match(/^[GCTA]*$/) == null) {
      throw new Error("Invalid input DNA.")
  }
  
  const table: {[key: string]: string} = {
    "G" : "C",
    "C" : "G",
    "T" : "A",
    "A" : "U" 
  }
  
  return rna.split("").map(x => table[x]).join("")

}