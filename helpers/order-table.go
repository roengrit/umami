package helpers

import (
	"bytes"
	"strconv"
	"strings"
	m "umami/models"
)

//HTMLTableTemplate _
const HTMLTableTemplate = `<tr>
							<td>{name}</td>
							<td>{qty}</td>
							<td>
								<div class="btn-group">
									{action}
								</div>
							</td>                             
						</tr>`

//HTMLTableActionEnable _
const HTMLTableActionEnable = `<a class="btn btn-primary " title="แก้ไข"  href="/table/?id={id}"><i class="fa fa-edit"></i></a>
							   <a class="btn btn-danger" title="ลบ" href="#"  onclick='confirmDeleteGlobal({id},"/table")'><i class="fa fa-trash-o"></i></a>`

//HTMLTableNotFoundRows _
const HTMLTableNotFoundRows = `<tr> <td  colspan="3" style="text-align:center;">*** ไม่พบข้อมูล ***</td></tr>`

//HTMLTableError _
const HTMLTableError = `<tr> <td  colspan="3" style="text-align:center;">{err}</td></tr>`

//GenTableHTML _
func GenTableHTML(lists []m.OrderTable) string {
	var hmtlBuffer bytes.Buffer
	for _, val := range lists {
		temp := strings.Replace(HTMLTableTemplate, "{name}", val.Name, -1)
		temp = strings.Replace(temp, "{qty}", strconv.Itoa(val.Qty), -1)
		tempAction := strings.Replace(HTMLTableActionEnable, "{id}", strconv.Itoa(val.ID), -1)
		temp = strings.Replace(temp, "{action}", tempAction, -1)
		hmtlBuffer.WriteString(temp)
	}
	return hmtlBuffer.String()
}
