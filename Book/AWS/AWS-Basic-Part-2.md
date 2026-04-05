 
###  หนังสือ  AWS จากภาคทฤษฎีไปภาคปฏิบัติ  
---
**📘 AWS จากภาคทฤษฎีไปภาคปฏิบัติ** 
**✍️ ผู้เขียน:** คงนคร จันทะคุณ  
**📅 อัปเดตล่าสุด:** เมษายน 2026  
**หมายเหตุ เนื้อหาในหนังสือ:**  
เนื้อหาในหนังสือ "AWS จากภาคทฤษฎีไปภาคปฏิบัติ" ใช้ AI ช่วยเขียน เพื่อทดสอบ AI Model ผู้เขียนเป็นผู้ออกแบบ ใช้ AI ช่วยจัดเรียง ซึ่งมีค่าใช้จ่ายพอสมควร ให้ใช้ฟรีก่อน ต้องการสนับสนุนเพื่อทำเนื้อหาแนวนี้ต่อ สามารถให้การสนับสนุนได้ครับ ตามกำลังศรัทธา 
📞 โทรศัพท์ / พร้อมเพย์: **0955088091**  

---

# 📘 สารบัญ (Table of Contents) เล่ม 2

**หนังสือ “AWS จากภาคทฤษฎีไปภาคปฏิบัติ”**  
*“AWS From Theory to Practice”*

| บทที่ | หัวข้อ | หน้า |
|-------|--------|------|
| 15 | AWS Certified Advanced Networking – Specialty (ANS) | 460 |
| 16 | AWS Certified Data Engineer – Associate (DEA) | 490 |
| 17 | AWS Certified DevOps Engineer – Professional (DOP) | 520 |
| ภาคผนวก | เทมเพลต โค้ดตัวอย่าง และเฉลยแบบฝึกหัด | 550 |

 
# 📘 บทที่ 15: AWS Certified Advanced Networking – Specialty (ANS-C01)  
## Chapter 15: AWS Certified Advanced Networking – Specialty (ANS-C01)  

---

## 🧱 โครงสร้างการทำงาน (Work Structure)  

**ไทย:**  
บทนี้เจาะลึกใบรับรอง AWS Certified Advanced Networking – Specialty (ANS-C01) สำหรับผู้ที่มีประสบการณ์ด้านเครือข่ายและต้องการพิสูจน์ทักษะการออกแบบ, implement, และจัดการโซลูชันเครือข่ายบน AWS ที่ซับซ้อน เนื้อหาครอบคลุมการออกแบบ VPC ขั้นสูง, การเชื่อมต่อ hybrid (VPN, Direct Connect), การทำ routing ข้ามภูมิภาค, การรักษาความปลอดภัยเครือข่าย, และการ optimize performance พร้อมตัวอย่างการใช้ AWS SDK สำหรับ Go ในการจัดการทรัพยากรเครือข่าย  

**English:**  
This chapter dives into the AWS Certified Advanced Networking – Specialty (ANS-C01) certification for those with networking experience who want to validate their skills in designing, implementing, and managing complex network solutions on AWS. It covers advanced VPC design, hybrid connectivity (VPN, Direct Connect), cross‑region routing, network security, and performance optimization, with examples of using the AWS SDK for Go to manage networking resources.  

---

## 🎯 วัตถุประสงค์แบบสั้นสำหรับทบทวน (Short Revision Objective)  

**ไทย:**  
เพื่อให้ผู้อ่านเข้าใจโครงสร้างข้อสอบ ANS-C01, เนื้อหาทั้ง 6 โดเมน, บริการเครือข่ายขั้นสูงของ AWS (Transit Gateway, Direct Connect, Route53 Resolver, Global Accelerator, VPC Lattice, etc.), และสามารถเตรียมตัวสอบ รวมถึงการเขียน Go เพื่อจัดการ VPC, Route53, และตรวจสอบเครือข่าย  

**English:**  
To enable readers to understand the ANS-C01 exam structure, six domains, advanced AWS networking services (Transit Gateway, Direct Connect, Route53 Resolver, Global Accelerator, VPC Lattice, etc.), and prepare effectively, including writing Go to manage VPC, Route53, and network monitoring.  

---

## 👥 กลุ่มเป้าหมาย (Target Audience)  

- Network Engineer / Architect ที่ทำงานบน AWS  
- Cloud Architect ที่ต้องออกแบบ hybrid และ multi‑region network  
- DevOps ที่ต้องจัดการ network infrastructure  
- ผู้ที่สอบผ่าน Solutions Architect Associate หรือ Professional แล้วต้องการต่อยอดด้านเครือข่าย  

---

## 📚 ความรู้พื้นฐาน (Prerequisites)  

- ความรู้เครือข่ายระดับกลางถึงสูง (OSI model, routing protocols (BGP), subnetting, VLAN, VPN, DNS)  
- ประสบการณ์ AWS อย่างน้อย 2 ปี รวมถึง VPC, Direct Connect, VPN, Transit Gateway  
- แนะนำให้สอบ Solutions Architect Associate หรือ Professional มาก่อน  

---

## 📝 เนื้อหาโดยย่อ (Abstract)  

**ไทย:**  
บทนี้สรุปข้อสอบ ANS-C01: จำนวนข้อ, เวลา, โดเมนหลัก 6 ด้าน (Design, Implementation, Security, Automation, Troubleshooting, Hybrid/Edge) พร้อมตัวอย่างคำถามระดับยาก, บริการที่ต้องรู้ลึก (VPC, VPN, Direct Connect, Transit Gateway, Route53, Global Accelerator, CloudFront, VPC Lattice, Network Firewall, etc.), และการเขียน Go เพื่อ automate การสร้าง VPC, Route53 records, และตรวจสอบ network metrics  

**English:**  
This chapter summarizes the ANS-C01 exam: number of questions, time, six domains (Design, Implementation, Security, Automation, Troubleshooting, Hybrid/Edge), sample difficult questions, essential services (VPC, VPN, Direct Connect, Transit Gateway, Route53, Global Accelerator, CloudFront, VPC Lattice, Network Firewall, etc.), and writing Go to automate VPC creation, Route53 records, and monitor network metrics.  

---

## 🔰 บทนำ (Introduction)  

**ไทย:**  
AWS Certified Advanced Networking – Specialty (ANS-C01) เป็นใบรับรองระดับ Specialty สำหรับผู้เชี่ยวชาญด้านเครือข่าย ข้อสอบจะทดสอบความสามารถในการออกแบบและ implement โซลูชันเครือข่ายที่ซับซ้อนบน AWS รวมถึงการเชื่อมต่อระหว่าง on‑premise และ AWS (hybrid), การทำ routing ข้าม region และข้าม VPC, การรักษาความปลอดภัยเครือข่าย (firewall, DDoS protection), และการ optimize performance สำหรับแอปพลิเคชันระดับโลก การสอบนี้เหมาะสำหรับ network engineer ที่มีประสบการณ์และต้องการพิสูจน์ความเชี่ยวชาญ  

**English:**  
The AWS Certified Advanced Networking – Specialty (ANS-C01) is a Specialty‑level certification for networking professionals. The exam tests the ability to design and implement complex network solutions on AWS, including hybrid connectivity (on‑premises to AWS), cross‑region and cross‑VPC routing, network security (firewalls, DDoS protection), and performance optimization for global applications. This exam suits experienced network engineers who want to validate their expertise.  

---

## 📖 บทนิยาม (Definitions)  

| คำศัพท์ (Term) | คำจำกัดความไทย (Thai Definition) | English Definition |
|----------------|----------------------------------|--------------------|
| VPC (Virtual Private Cloud) | เครือข่ายส่วนตัวใน AWS | Isolated network in AWS. |
| Subnet | ช่วงย่อยของ IP ภายใน VPC (public หรือ private) | IP range subdivision within a VPC (public or private). |
| Route Table | ตารางกำหนดเส้นทางของ traffic ใน VPC | Table defining traffic routing within a VPC. |
| Internet Gateway (IGW) | gateway สำหรับเชื่อมต่อ VPC กับ internet | Gateway connecting VPC to the internet. |
| NAT Gateway | ให้ instances ใน private subnet ออก internet (แต่ภายนอกเข้าหาไม่ได้) | Allows private subnet instances to access internet (but not inbound). |
| VPN (Virtual Private Network) | การเชื่อมต่อ encrypted ผ่าน internet ระหว่าง on‑premise กับ AWS | Encrypted connection over internet between on‑premises and AWS. |
| Direct Connect (DX) | การเชื่อมต่อ dedicated, private ระหว่าง on‑premise กับ AWS | Dedicated, private connection between on‑premises and AWS. |
| Transit Gateway (TGW) | Hub สำหรับเชื่อมต่อ VPC, VPN, DX หลายๆ จุด | Hub connecting multiple VPCs, VPNs, DX. |
| VPC Peering | การเชื่อมต่อตรงระหว่างสอง VPC (ไม่ผ่าน TGW) | Direct connection between two VPCs (no TGW). |
| Route53 Resolver | ให้ DNS resolution สำหรับ hybrid network (on‑premise ↔ AWS) | DNS resolution for hybrid networks. |
| Global Accelerator | บริการ加速 traffic ทั่วโลกด้วย Anycast IP | Global traffic acceleration using Anycast IP. |
| VPC Lattice | บริการเชื่อมต่อ service ข้าม VPC และ account แบบ application‑layer | Cross‑VPC/cross‑account service connectivity at application layer. |
| Network Firewall | Managed firewall สำหรับ VPC (stateful, 规则-based) | Managed firewall for VPC (stateful, rule‑based). |

---

## 🔧 ANS-C01 คืออะไร? มีเนื้อหาอะไรบ้าง?  

### 1. ANS-C01 คืออะไร  
**ไทย:**  
ANS-C01 คือรหัสข้อสอบ AWS Certified Advanced Networking – Specialty (อัปเดตล่าสุด 2022-2023) ทดสอบความสามารถในการออกแบบ, implement, และ troubleshoot โซลูชันเครือข่ายบน AWS ที่ซับซ้อน โดยเฉพาะ hybrid networking, routing ขั้นสูง, และ security  

**English:**  
ANS-C01 is the exam code for AWS Certified Advanced Networking – Specialty (latest update 2022-2023). It tests the ability to design, implement, and troubleshoot complex network solutions on AWS, especially hybrid networking, advanced routing, and security.  

### 2. เนื้อหาข้อสอบแบ่งเป็น 6 โดเมน (Domains)  

| โดเมน (Domain) | น้ำหนัก (Weight) | หัวข้อหลัก (Key topics) |
|----------------|------------------|--------------------------|
| Network Design | 22% | VPC design (CIDR, subnet sizing, multi‑AZ), hybrid connectivity (DX, VPN), routing policies (static, BGP), high availability, disaster recovery |
| Network Implementation | 18% | การสร้าง VPC, subnets, route tables, IGW, NAT, VPC peering, Transit Gateway, VPN, Direct Connect, Global Accelerator, VPC Lattice |
| Network Security | 20% | Security groups, NACLs, AWS WAF, AWS Shield, Network Firewall, encryption in transit (TLS, IPsec), VPC endpoints (Gateway, Interface) |
| Network Automation | 12% | Infrastructure as Code (CloudFormation, Terraform), AWS CLI, SDK (Go, Python), AWS Config rules, VPC Flow Logs automation |
| Network Troubleshooting | 14% | การวิเคราะห์ VPC Flow Logs, Route Analyzer, Reachability Analyzer, CloudWatch metrics, X-Ray, packet capture, BGP troubleshooting |
| Hybrid & Edge Networking | 14% | Direct Connect (public, private, transit VIF), VPN CloudHub, Transit Gateway + DX, Route53 Resolver (inbound/outbound endpoints), AWS Outposts, Local Zones |

### 3. รูปแบบข้อสอบ  

| รายการ | รายละเอียด |
|--------|-------------|
| จำนวนข้อ | 65 (รวม 15 ข้อที่ไม่นับคะแนน) |
| เวลา | 170 นาที (2 ชั่วโมง 50 นาที) |
| รูปแบบ | Multiple choice, multiple answer |
| คะแนนผ่าน | 750/1000 (75%) |
| ค่าสอบ | 300 USD |
| ภาษา | อังกฤษ, ญี่ปุ่น, เกาหลี, จีน |

### 4. บริการที่ต้องรู้ลึกสำหรับ ANS  

| บริการ | ความสำคัญ | หัวข้อที่ต้องรู้ |
|--------|------------|----------------|
| VPC | สูงมาก | CIDR, subnets, route tables, IGW, NAT, endpoints, peering |
| Transit Gateway | สูงมาก | attachments, route tables, cross‑region peering, multicast, appliance mode |
| Direct Connect | สูงมาก | public/private/transit VIF, LAG, MACsec, BGP, DX Gateway |
| VPN (Site‑to‑Site) | สูง | VPN CloudHub, dynamic routing (BGP), tunnel options, acceleration |
| Route53 | สูง | Resolver (inbound/outbound), private hosted zones, routing policies (failover, latency, geolocation, weighted) |
| Global Accelerator | ปานกลาง | anycast IP, endpoint groups, health checks, traffic dial |
| VPC Lattice | ใหม่ (ปานกลาง) | service network, target groups, auth policies |
| Network Firewall | ปานกลาง | stateful rule groups, suricata format, logging |
| CloudFront | ปานกลาง | origin shield, custom headers, lambda@edge, field‑level encryption |
| VPC Flow Logs | สูง | log format, query with Athena, troubleshooting |

### 5. ตัวอย่างคำถาม (Sample Question – ระดับยาก)  

**คำถาม:** บริษัทมี on‑premise data center 2 แห่ง (ในประเทศ A และ B) และ AWS region 3 แห่ง (us-east-1, eu-west-1, ap-southeast-1) ต้องการเชื่อมต่อทุก site ด้วยความหน่วงต่ำและมี failover อัตโนมัติ ควรออกแบบอย่างไร  

A. สร้าง Direct Connect จากแต่ละ data center ไปยัง region ที่ใกล้ที่สุด แล้วใช้ Transit Gateway เชื่อมต่อระหว่าง region  
B. สร้าง VPN tunnels ระหว่าง data center ทั้งหมด  
C. ใช้ VPC peering เชื่อมต่อทุก VPC และทำ VPN จาก data center ไปยัง VPC หลัก  
D. ใช้ Global Accelerator สำหรับทุกการเชื่อมต่อ  

**เฉลย:** A (Direct Connect + Transit Gateway ให้ความหน่วงต่ำและ routing แบบ mesh ผ่าน TGW, cross‑region peering ของ TGW เชื่อมภูมิภาค)  

---

## 🔄 ออกแบบ Workflow (Workflow Design)  

### ภาพรวม: Hybrid Network ด้วย Direct Connect + Transit Gateway  

**ไทย:**  
On‑premise router → Direct Connect (private VIF) → Direct Connect Gateway → Transit Gateway → VPC attachments (หลาย VPC) → ภายใน AWS อาจมี TGW peering ข้าม region  

**English:**  
On‑premises router → Direct Connect (private VIF) → Direct Connect Gateway → Transit Gateway → VPC attachments (multiple VPCs) → optionally TGW peering across regions.  

### Mermaid Flowchart  

```mermaid
flowchart TB
    OnPrem[On-Premise Router] --> DX[Direct Connect]
    DX --> DXG[Direct Connect Gateway]
    DXG --> TGW[Transit Gateway]
    TGW --> VPC1[VPC A - Production]
    TGW --> VPC2[VPC B - Shared Services]
    TGW --> VPC3[VPC C - Dev/Test]
    TGW -.-> TGW2[Transit Gateway Peering<br/>(cross-region)]
    TGW2 --> VPC4[VPC D - DR Region]
```

### คำอธิบายแบบละเอียด (Detailed Explanation)  

| ขั้นตอน | คำอธิบาย (ไทย) | Explanation (English) |
|---------|----------------|------------------------|
| 1 | On‑premise router เชื่อมต่อ AWS ผ่าน Direct Connect (private VIF) | On‑premises router connects via Direct Connect private VIF. |
| 2 | Direct Connect Gateway (DX Gateway) ทำหน้าที่รวม DX VIF หลายอันเข้าด้วยกัน และเชื่อมต่อไปยัง Transit Gateway | DX Gateway aggregates multiple DX VIFs and connects to Transit Gateway. |
| 3 | Transit Gateway (TGW) ทำหน้าที่เป็น hub เชื่อมต่อ VPC หลายตัวเข้าด้วยกัน | TGW acts as hub connecting multiple VPCs. |
| 4 | TGW สามารถ peering กับ TGW ในอีก region เพื่อขยาย connectivity ข้ามภูมิภาค | TGW can peer with another TGW in a different region for cross‑region connectivity. |
| 5 | VPC ต่างๆ สามารถสื่อสารกันผ่าน TGW โดยไม่ต้องมี VPC peering แบบ一对一 | VPCs communicate via TGW without need for one‑to‑one peering. |

---

## 💻 ตัวอย่างโค้ดที่รันได้จริง (Runnable Code Example)  

### 1. การสร้าง VPC และ Subnet ด้วย Go SDK  

```go
// create_vpc.go
// สร้าง VPC, subnet, internet gateway, route table ด้วย Go SDK
// Create VPC, subnet, internet gateway, route table using Go SDK

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}
	client := ec2.NewFromConfig(cfg)

	// 1. สร้าง VPC (CIDR 10.0.0.0/16)
	vpcResp, err := client.CreateVpc(context.TODO(), &ec2.CreateVpcInput{
		CidrBlock: stringPtr("10.0.0.0/16"),
	})
	if err != nil {
		log.Fatal(err)
	}
	vpcId := *vpcResp.Vpc.VpcId
	fmt.Printf("Created VPC: %s\n", vpcId)

	// 2. สร้าง Internet Gateway
	igwResp, err := client.CreateInternetGateway(context.TODO(), &ec2.CreateInternetGatewayInput{})
	if err != nil {
		log.Fatal(err)
	}
	igwId := *igwResp.InternetGateway.InternetGatewayId
	fmt.Printf("Created IGW: %s\n", igwId)

	// 3. Attach IGW กับ VPC
	_, err = client.AttachInternetGateway(context.TODO(), &ec2.AttachInternetGatewayInput{
		InternetGatewayId: &igwId,
		VpcId:             &vpcId,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 4. สร้าง Subnet (public)
	subnetResp, err := client.CreateSubnet(context.TODO(), &ec2.CreateSubnetInput{
		VpcId:            &vpcId,
		CidrBlock:        stringPtr("10.0.1.0/24"),
		AvailabilityZone: stringPtr("us-east-1a"),
	})
	if err != nil {
		log.Fatal(err)
	}
	subnetId := *subnetResp.Subnet.SubnetId
	fmt.Printf("Created Subnet: %s\n", subnetId)

	// 5. สร้าง Route Table และ route ไป IGW
	rtResp, err := client.CreateRouteTable(context.TODO(), &ec2.CreateRouteTableInput{
		VpcId: &vpcId,
	})
	if err != nil {
		log.Fatal(err)
	}
	rtId := *rtResp.RouteTable.RouteTableId

	_, err = client.CreateRoute(context.TODO(), &ec2.CreateRouteInput{
		RouteTableId:         &rtId,
		DestinationCidrBlock: stringPtr("0.0.0.0/0"),
		GatewayId:            &igwId,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 6. Associate route table กับ subnet
	_, err = client.AssociateRouteTable(context.TODO(), &ec2.AssociateRouteTableInput{
		RouteTableId: &rtId,
		SubnetId:     &subnetId,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("VPC setup complete!")
}

func stringPtr(s string) *string { return &s }
```

### 2. การสร้าง Route53 Private Hosted Zone และบันทึก  

```go
// route53_private_zone.go
// สร้าง private hosted zone ใน VPC และเพิ่ม A record
// Create private hosted zone within VPC and add A record

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	client := route53.NewFromConfig(cfg)

	vpcId := "vpc-12345678"   // VPC ID ที่ต้องการ关联
	zoneName := "internal.example.com"

	// สร้าง private hosted zone
	resp, err := client.CreateHostedZone(context.TODO(), &route53.CreateHostedZoneInput{
		Name: &zoneName,
		VPC: &types.VPC{
			VPCId:     &vpcId,
			VPCRegion: types.VPCRegionUsEast1,
		},
		CallerReference: stringPtr("ref-12345"),
	})
	if err != nil {
		log.Fatal(err)
	}
	zoneId := *resp.HostedZone.Id
	fmt.Printf("Created private hosted zone: %s (ID: %s)\n", zoneName, zoneId)

	// เพิ่ม A record
	recordName := "api.internal.example.com"
	recordValue := "10.0.1.100"

	_, err = client.ChangeResourceRecordSets(context.TODO(), &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: &zoneId,
		ChangeBatch: &types.ChangeBatch{
			Changes: []types.Change{
				{
					Action: types.ChangeActionCreate,
					ResourceRecordSet: &types.ResourceRecordSet{
						Name: &recordName,
						Type: types.RRTypeA,
						TTL:  int64Ptr(300),
						ResourceRecords: []types.ResourceRecord{
							{Value: &recordValue},
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Added A record: %s -> %s\n", recordName, recordValue)
}

func int64Ptr(i int64) *int64 { return &i }
```

### 3. การเปิด VPC Flow Logs และส่งไปยัง CloudWatch  

```go
// vpc_flow_logs.go
// เปิด VPC Flow Logs สำหรับ VPC และส่ง logs ไปยัง CloudWatch Logs
// Enable VPC Flow Logs for a VPC and send logs to CloudWatch Logs

func enableFlowLogs(client *ec2.Client, vpcId, logGroupName string) error {
	_, err := client.CreateFlowLogs(context.TODO(), &ec2.CreateFlowLogsInput{
		ResourceIds: []string{vpcId},
		ResourceType: types.FlowLogsResourceTypeVpc,
		TrafficType:  types.TrafficTypeAll,
		LogDestinationType: types.LogDestinationTypeCloudWatchLogs,
		LogGroupName: &logGroupName,
		DeliverLogsPermissionArn: stringPtr("arn:aws:iam::ACCOUNT:role/FlowLogsRole"),
	})
	return err
}
```

---

## 📌 กรณีศึกษาและแนวทางแก้ไขปัญหา (Case Study & Troubleshooting)  

### กรณีศึกษา: การ troubleshoot connectivity ระหว่าง on‑premise และ VPC  

**ปัญหา:** หลังตั้งค่า Direct Connect และ VPN backup แล้ว traffic บางส่วนไม่สามารถเข้าถึง VPC ได้  
**แนวทางแก้ไข (ตาม ANS):**  
- ตรวจสอบ route tables บน TGW และ VPC ว่า route ไปยัง DX หรือ VPN ถูกต้อง  
- ใช้ Reachability Analyzer ตรวจสอบเส้นทางจาก on‑premise IP ไปยัง destination IP ใน VPC  
- ตรวจสอบ BGP session (ถ้าใช้ dynamic routing) ว่า received routes ถูกต้อง  
- ดู VPC Flow Logs ว่า packets ถูก drop ที่ security group หรือ NACL หรือไม่  
**ผลลัพธ์:** พบว่า route table บน TGW ไม่มี route กลับไปยัง DX Gateway, เพิ่ม route แล้วแก้ไขได้  

### ปัญหาที่พบบ่อยในการสอบ ANS  

| ปัญหา (Issue) | สาเหตุ (Cause) | วิธีแก้ไข (Solution) |
|----------------|----------------|----------------------|
| งบประมาณสูง | ใช้ Direct Connect ไม่ถูกต้อง | ใช้ VPN เป็น backup หรือใช้ DX ในโหมด "hosted" สำหรับปริมาณน้อย |
| routing loop | TGW route table configuration ผิด | ใช้ static routes หรือ BGP attributes (local preference, AS path) เพื่อควบคุมเส้นทาง |
| DNS resolution ล้มเหลวใน hybrid | ไม่ได้ตั้ง Route53 Resolver | ตั้ง inbound/outbound endpoints เพื่อ forward queries ระหว่าง on‑premise DNS และ AWS |
| รู้สึกว่า VPC peering กับ TGW ต่างกันอย่างไร | VPC peering = 一对一, ไม่ support transitive; TGW = hub-and-spoke, support transitive | จดจำ: ใช้ TGW เมื่อมี VPC มากกว่า 2-3 ตัว |
| การรักษาความปลอดภัย | ใช้ security group และ NACL ไม่ถูกต้อง | Security group (stateful) ที่ instance level; NACL (stateless) ที่ subnet level |

---

## 📁 เทมเพลตและตัวอย่างเพิ่มเติม  

### แผนการเตรียมตัว 12 สัปดาห์ (ANS-C01)  

| สัปดาห์ | กิจกรรม |
|---------|---------|
| 1 | ทบทวน networking fundamentals (OSI, TCP/IP, routing, BGP, DNS) |
| 2 | VPC ขั้นสูง: CIDR, subnet, route tables, IGW, NAT, endpoints, peering |
| 3 | Transit Gateway: attachments, route tables, cross‑region peering |
| 4 | Direct Connect: public/private/transit VIF, LAG, MACsec, DX Gateway |
| 5 | VPN: site‑to‑site, VPN CloudHub, BGP, acceleration |
| 6 | Route53: private hosted zones, resolver, routing policies |
| 7 | Global Accelerator, CloudFront, VPC Lattice |
| 8 | Network Security: Security groups, NACL, WAF, Shield, Network Firewall, encryption |
| 9 | Automation: CloudFormation, CDK, CLI, SDK, Config rules |
| 10 | Monitoring & Troubleshooting: VPC Flow Logs, Reachability Analyzer, CloudWatch, X-Ray |
| 11 | Hybrid & Edge: Outposts, Local Zones, IoT Core networking |
| 12 | ทำ practice exam 3-5 ชุด, ทบทวนจุดอ่อน, สอบจริง |

### Checklist ก่อนสอบ ANS  

- [ ] รู้ CIDR calculation และ subnet sizing  
- [ ] รู้ความแตกต่างระหว่าง VPC Peering และ Transit Gateway  
- [ ] รู้วิธีตั้ง Direct Connect (public/private/transit VIF) และ DX Gateway  
- [ ] รู้ BGP basics (ASN, peering, route advertisement, local preference, MED)  
- [ ] รู้วิธีทำ VPN แบบ static และ dynamic routing  
- [ ] รู้ Route53 resolver inbound/outbound  
- [ ] รู้วิธีใช้ VPC Flow Logs และ Athena วิเคราะห์  
- [ ] รู้ Reachability Analyzer และ Network Manager  
- [ ] รู้ AWS Network Firewall rule groups  
- [ ] รู้ Global Accelerator กับ CloudFront ต่างกันอย่างไร  

---

## 📊 ตารางเปรียบเทียบ VPC Peering vs Transit Gateway  

| คุณสมบัติ | VPC Peering | Transit Gateway |
|-----------|-------------|------------------|
| Topology | point‑to‑point | hub‑and‑spoke |
| Transitive routing | ไม่ | ใช่ (ผ่าน TGW) |
| ข้ามบัญชี | ใช่ | ใช่ (ผ่าน RAM) |
| ข้าม region | ใช่ (inter-region peering) | ใช่ (TGW peering) |
| การจัดการ route | manual หรือ auto (ไม่ซับซ้อน) | route tables แยกต่อ attachment |
| Scaling | limit 125 peerings ต่อ VPC | รองรับ 5000 attachments ต่อ TGW |
| ราคา | ถูกกว่าสำหรับการเชื่อมต่อน้อย | คุ้มกว่าสำหรับการเชื่อมต่อมาก |

---

## 📝 สรุป (Summary)  

### ✅ ประโยชน์ที่ได้รับ (Benefits)  
- พิสูจน์ทักษะ networking ขั้นสูงบน AWS  
- ออกแบบ hybrid network ที่ปลอดภัยและ scalable  
- เงินเดือนและตำแหน่งสูงขึ้นสำหรับ network specialist  
- สามารถ troubleshoot ปัญหาเครือข่ายที่ซับซ้อนได้  

### ⚠️ ข้อควรระวัง (Cautions)  
- ต้องมีความรู้ networking มาก่อน (ไม่เหมาะสำหรับผู้เริ่มต้น)  
- ค่าสอบแพง (300 USD)  
- เนื้อหาลึกและกว้าง ต้องใช้เวลาเตรียม 2-3 เดือน  

### 👍 ข้อดี (Advantages)  
- เป็นใบรับรองที่หายากและมีค่า  
- ครอบคลุม hybrid และ modern networking (VPC Lattice)  
- ใช้ automation กับ SDK ได้จริง  

### 👎 ข้อเสีย (Disadvantages)  
- ต้องอัปเดตความรู้บริการใหม่ ๆ (เช่น VPC Lattice)  
- ข้อสอบยาก อัตราการผ่านต่ำ  
- ไม่มี lab ปฏิบัติ (ต่างจาก Professional บางตัว)  

### 🚫 ข้อห้าม (Prohibitions)  
- ห้ามใช้ default VPC สำหรับ production (ต้องออกแบบเอง)  
- ห้ามเปิด security group หรือ NACL แบบกว้างเกินไป (0.0.0.0/0) โดยไม่จำเป็น  
- ห้ามใช้ VPC peering แทน TGW เมื่อมีการเชื่อมต่อมากกว่า 5 VPC  

---

## 🧩 แบบฝึกหัดท้ายบท (Exercises)  

**ข้อ 1:** ANS-C01 มีน้ำหนักของโดเมน "Network Design" กี่เปอร์เซ็นต์  
**ข้อ 2:** ความแตกต่างหลักระหว่าง Internet Gateway กับ NAT Gateway คืออะไร  
**ข้อ 3:** หากต้องการเชื่อมต่อ VPC จำนวน 10 VPC ใน region เดียวกันด้วยต้นทุนการดูแลที่ต่ำ ควรใช้บริการใด  
**ข้อ 4:** Direct Connect Gateway (DX Gateway) มีหน้าที่อะไร  
**ข้อ 5:** Route53 Resolver inbound endpoint ใช้ทำอะไร  
**ข้อ 6:** BGP ในบริบทของ AWS VPN ใช้ทำอะไร  
**ข้อ 7:** ข้อสอบ ANS-C01 มีจำนวนข้อและเวลาเท่าไร  
**ข้อ 8:** จงเขียน Go code เพื่อสร้าง Route53 A record ใน public hosted zone  
**ข้อ 9:** AWS Network Firewall แตกต่างจาก Security Group อย่างไร  
**ข้อ 10:** หากต้องการวิเคราะห์ VPC Flow Logs ด้วย SQL ควรใช้บริการใด  

---

## 🔐 เฉลยแบบฝึกหัด (Answer Key)  

**ข้อ 1:** 22%  
**ข้อ 2:** IGW ให้ inbound/outbound internet access สำหรับ public subnet; NAT Gateway ให้ outbound internet access สำหรับ private subnet (ไม่ให้ inbound)  
**ข้อ 3:** Transit Gateway (TGW)  
**ข้อ 4:** เชื่อมต่อ Direct Connect VIF หนึ่งตัวหรือหลายตัว เข้ากับ Transit Gateway หรือ VPC (ผ่าน VPC attachment)  
**ข้อ 5:** ให้ DNS queries จาก on‑premise มาที่ AWS Route53 private hosted zones ได้  
**ข้อ 6:** แลกเปลี่ยน route ระหว่าง AWS VPN endpoint และ on‑premise router แบบไดนามิก (dynamic routing)  
**ข้อ 7:** 65 ข้อ, 170 นาที  
**ข้อ 8:**  
```go
func createARecord(client *route53.Client, zoneId, name, value string) error {
    _, err := client.ChangeResourceRecordSets(context.TODO(), &route53.ChangeResourceRecordSetsInput{
        HostedZoneId: &zoneId,
        ChangeBatch: &types.ChangeBatch{
            Changes: []types.Change{{
                Action: types.ChangeActionUpsert,
                ResourceRecordSet: &types.ResourceRecordSet{
                    Name: &name,
                    Type: types.RRTypeA,
                    TTL: int64Ptr(300),
                    ResourceRecords: []types.ResourceRecord{{Value: &value}},
                },
            }},
        },
    })
    return err
}
```  
**ข้อ 9:** Security Group ทำงานที่ instance level, stateful; Network Firewall ทำงานที่ VPC level, stateful, รองรับ rule sets แบบ Suricata, ให้ centralized inspection  
**ข้อ 10:** Amazon Athena (query logs ที่เก็บใน S3)  

---

## 📚 แหล่งอ้างอิง (References)  

1. AWS Official Exam Guide – ANS-C01  
2. AWS VPC and Networking Documentation  
3. AWS Transit Gateway Guide  
4. Direct Connect User Guide  
5. TutorialsDojo – ANS-C01 Practice Exams  
6. AWS Networking Workshops (GitHub)  

---

**✍️ ผู้เขียน:** คงนคร จันทะคุณ  
**📅 อัปเดตล่าสุด:** เมษายน 2026  

**หมายเหตุ เนื้อหาในหนังสือ:**  
เนื้อหาในหนังสือ "AWS จากภาคทฤษฎีไปภาคปฏิบัติ" ใช้ AI ช่วยเขียน เพื่อทดสอบ AI Model ผู้เขียนเป็นผู้ออกแบบ ใช้ AI ช่วยจัดเรียง ซึ่งมีค่าใช้จ่ายพอสมควร ให้ใช้ฟรีก่อน ต้องการสนับสนุนเพื่อทำเนื้อหาแนวนี้ต่อ สามารถให้การสนับสนุนได้ครับ ตามกำลังศรัทธา  
📞 โทรศัพท์ / พร้อมเพย์: **0955088091**