package main

import (
	"fmt"
	"strings"
)

type Transformation struct{
	from,to string
}

func processTransformations() []Transformation {
	data :=`Al => ThF
Al => ThRnFAr
B => BCa
B => TiB
B => TiRnFAr
Ca => CaCa
Ca => PB
Ca => PRnFAr
Ca => SiRnFYFAr
Ca => SiRnMgAr
Ca => SiTh
F => CaF
F => PMg
F => SiAl
H => CRnAlAr
H => CRnFYFYFAr
H => CRnFYMgAr
H => CRnMgYFAr
H => HCa
H => NRnFYFAr
H => NRnMgAr
H => NTh
H => OB
H => ORnFAr
Mg => BF
Mg => TiMg
N => CRnFAr
N => HSi
O => CRnFYFAr
O => CRnMgAr
O => HP
O => NRnFAr
O => OTi
P => CaP
P => PTi
P => SiRnFAr
Si => CaSi
Th => ThCa
Ti => BP
Ti => TiTi
e => HF
e => NAl
e => OMg`
	rows := strings.Split(data,"\n");
	ts := []Transformation{}
	for _,r := range rows {
		r = strings.TrimSpace(r)
		e := strings.Split(r," => ")
		ts = append(ts, Transformation{e[0],e[1]})
	}
	return ts
}
func main() {
	masterMolecule := "CRnSiRnCaPTiMgYCaPTiRnFArSiThFArCaSiThSiThPBCaCaSiRnSiRnTiTiMgArPBCaPMgYPTiRnFArFArCaSiRnBPMgArPRnCaPTiRnFArCaSiThCaCaFArPBCaCaPTiTiRnFArCaSiRnSiAlYSiThRnFArArCaSiRnBFArCaCaSiRnSiThCaCaCaFYCaPTiBCaSiThCaSiThPMgArSiRnCaPBFYCaCaFArCaCaCaCaSiThCaSiRnPRnFArPBSiThPRnFArSiRnMgArCaFYFArCaSiRnSiAlArTiTiTiTiTiTiTiRnPMgArPTiTiTiBSiRnSiAlArTiTiRnPMgArCaFYBPBPTiRnSiRnMgArSiThCaFArCaSiThFArPRnFArCaSiRnTiBSiThSiRnSiAlYCaFArPRnFArSiThCaFArCaCaSiThCaCaCaSiRnPRnCaFArFYPMgArCaPBCaPBSiRnFYPBCaFArCaSiAl"
	transformations := processTransformations()
	found := map[string]bool{}
	for i,l := range masterMolecule {
		for _,t := range transformations {
			var check string
			if len(t.from)==1{
				check = string(l)
			}else if len(t.from)==2 && i<len(masterMolecule)-1{
				check = string(masterMolecule[i])+string(masterMolecule[i+1])
			}
			if t.from==check{
				temp := string(masterMolecule[:i])+t.to
				if len(t.from)==2{
					temp += string(masterMolecule[i+2:])
				}else{
					temp += string(masterMolecule[i+1:])
				}
				found[temp]=true
			}
		}
	}
	fmt.Println(len(found))

}