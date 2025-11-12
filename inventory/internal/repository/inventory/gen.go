package inventory

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
	"github.com/mercuryqa/rocket-lab/inventory/internal/repository/converter"
	repoModel "github.com/mercuryqa/rocket-lab/inventory/internal/repository/model"
	pb "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

func DimensionsProtoToModel(pbDim *pb.Dimensions) *model.Dimensions {
	if pbDim == nil {
		return &model.Dimensions{}
	}
	return &model.Dimensions{
		Length: pbDim.Length,
		Width:  pbDim.Width,
		Height: pbDim.Height,
		Weight: pbDim.Weight,
	}
}

func ManufacturerProtoToModel(pbMan *pb.Manufacturer) *model.Manufacturer {
	if pbMan == nil {
		return &model.Manufacturer{}
	}
	return &model.Manufacturer{
		Name:    pbMan.Name,
		Country: pbMan.Country,
	}
}

func PartProtoToModel(pbPart *pb.Part) *model.Part {
	return &model.Part{
		UUID:          pbPart.Uuid,
		Name:          pbPart.Name,
		Description:   pbPart.Description,
		Price:         pbPart.Price,
		StockQuantity: pbPart.StockQuantity,
		Category:      model.Category(pbPart.Category),
		Dimensions:    *DimensionsProtoToModel(pbPart.Dimensions),     // разыменовываем
		Manufacturer:  *ManufacturerProtoToModel(pbPart.Manufacturer), // разыменовываем
		Tags:          pbPart.Tags,
	}
}

func GenerateSampleData(r *InventoryRepository) {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	ts := timestamppb.New(now)

	parts := []*pb.Part{
		{
			Uuid:          "1",
			Name:          "Wing Alpha",
			Description:   "Lightweight wing",
			Price:         45_000.0,
			StockQuantity: 50,
			Category:      pb.Category_WING,
			Dimensions:    &pb.Dimensions{Length: 150.0, Width: 20.0, Height: 10.0, Weight: 120.0},
			Manufacturer:  &pb.Manufacturer{Name: "EuroAero", Country: "Germany", Website: "https://euroaero.example"},
			Tags:          []string{"aero", "wing"},
			Metadata: map[string]*pb.Value{
				"material": {Kind: &pb.Value_StringValue{StringValue: "carbon"}},
				"tested":   {Kind: &pb.Value_BoolValue{BoolValue: true}},
			},
			CreatedAt: ts,
			UpdatedAt: ts,
		},
		{
			Uuid:          "2",
			Name:          "Main Engine",
			Description:   "High-thrust main engine",
			Price:         1_200_000.0,
			StockQuantity: 5,
			Category:      pb.Category_ENGINE,
			Dimensions:    &pb.Dimensions{Length: 300.0, Width: 200.0, Height: 400.0, Weight: 1200.0},
			Manufacturer:  &pb.Manufacturer{Name: "AstroWorks", Country: "USA", Website: "https://astro.example"},
			Tags:          []string{"engine", "core"},
			Metadata: map[string]*pb.Value{
				"max_thrust": {Kind: &pb.Value_DoubleValue{DoubleValue: 7600.0}},
				"fuel_type":  {Kind: &pb.Value_StringValue{StringValue: "liquid hydrogen"}},
			},
			CreatedAt: ts,
			UpdatedAt: ts,
		},
		{
			Uuid:          "3",
			Name:          "Fuel Tank A1",
			Description:   "Cryogenic fuel storage tank",
			Price:         85_000.0,
			StockQuantity: 20,
			Category:      pb.Category_FUEL,
			Dimensions:    &pb.Dimensions{Length: 200.0, Width: 80.0, Height: 80.0, Weight: 300.0},
			Manufacturer:  &pb.Manufacturer{Name: "Orbital Tanks", Country: "USA", Website: "https://orbitaltanks.example"},
			Tags:          []string{"fuel", "storage"},
			Metadata: map[string]*pb.Value{
				"capacity_l": {Kind: &pb.Value_DoubleValue{DoubleValue: 5000.0}},
				"tested":     {Kind: &pb.Value_BoolValue{BoolValue: true}},
			},
			CreatedAt: ts,
			UpdatedAt: ts,
		},
		{
			Uuid:          "4",
			Name:          "Porthole Glass X2",
			Description:   "Reinforced transparent porthole",
			Price:         12_000.0,
			StockQuantity: 15,
			Category:      pb.Category_PORTHOLE,
			Dimensions:    &pb.Dimensions{Length: 50.0, Width: 50.0, Height: 10.0, Weight: 50.0},
			Manufacturer:  &pb.Manufacturer{Name: "SpaceVision", Country: "Japan", Website: "https://spacevision.example"},
			Tags:          []string{"glass", "safety"},
			Metadata: map[string]*pb.Value{
				"material": {Kind: &pb.Value_StringValue{StringValue: "reinforced glass"}},
				"tested":   {Kind: &pb.Value_BoolValue{BoolValue: true}},
			},
			CreatedAt: ts,
			UpdatedAt: ts,
		},
		{
			Uuid:          "5",
			Name:          "Aux Engine Booster",
			Description:   "Secondary engine for maneuvering",
			Price:         350_000.0,
			StockQuantity: 10,
			Category:      pb.Category_ENGINE,
			Dimensions:    &pb.Dimensions{Length: 150.0, Width: 100.0, Height: 150.0, Weight: 600.0},
			Manufacturer:  &pb.Manufacturer{Name: "AstroWorks", Country: "USA", Website: "https://astro.example"},
			Tags:          []string{"engine", "booster"},
			Metadata: map[string]*pb.Value{
				"max_thrust": {Kind: &pb.Value_DoubleValue{DoubleValue: 3200.0}},
			},
			CreatedAt: ts,
			UpdatedAt: ts,
		},
		{
			Uuid:          "6",
			Name:          "Wing Beta",
			Description:   "Secondary wing module",
			Price:         40_000.0,
			StockQuantity: 40,
			Category:      pb.Category_WING,
			Dimensions:    &pb.Dimensions{Length: 140.0, Width: 18.0, Height: 12.0, Weight: 110.0},
			Manufacturer:  &pb.Manufacturer{Name: "EuroAero", Country: "Germany", Website: "https://euroaero.example"},
			Tags:          []string{"wing", "backup"},
			Metadata: map[string]*pb.Value{
				"material": {Kind: &pb.Value_StringValue{StringValue: "alloy"}},
				"tested":   {Kind: &pb.Value_BoolValue{BoolValue: true}},
			},
			CreatedAt: ts,
			UpdatedAt: ts,
		},
		{
			Uuid:          "7",
			Name:          "Fuel Pump X",
			Description:   "Fuel transfer pump",
			Price:         15_000.0,
			StockQuantity: 25,
			Category:      pb.Category_FUEL,
			Dimensions:    &pb.Dimensions{Length: 50.0, Width: 30.0, Height: 30.0, Weight: 25.0},
			Manufacturer:  &pb.Manufacturer{Name: "Orbital Tanks", Country: "Germany", Website: "https://orbitaltanks.example"},
			Tags:          []string{"fuel", "pump"},
			Metadata: map[string]*pb.Value{
				"flow_rate_l_per_min": {Kind: &pb.Value_DoubleValue{DoubleValue: 250.0}},
			},
			CreatedAt: ts,
			UpdatedAt: ts,
		},
		{
			Uuid:          "8",
			Name:          "Porthole Glass Y3",
			Description:   "Extra-large reinforced porthole",
			Price:         18_000.0,
			StockQuantity: 8,
			Category:      pb.Category_PORTHOLE,
			Dimensions:    &pb.Dimensions{Length: 70.0, Width: 70.0, Height: 10.0, Weight: 70.0},
			Manufacturer:  &pb.Manufacturer{Name: "SpaceVision", Country: "Japan", Website: "https://spacevision.example"},
			Tags:          []string{"glass", "reinforced"},
			Metadata: map[string]*pb.Value{
				"material": {Kind: &pb.Value_StringValue{StringValue: "reinforced glass"}},
				"tested":   {Kind: &pb.Value_BoolValue{BoolValue: true}},
			},
			CreatedAt: ts,
			UpdatedAt: ts,
		},
		{
			Uuid:          "9",
			Name:          "Navigation CPU",
			Description:   "Central navigation computer",
			Price:         250_000.0,
			StockQuantity: 12,
			Category:      pb.Category_UNKNOWN,
			Dimensions:    &pb.Dimensions{Length: 60.0, Width: 40.0, Height: 20.0, Weight: 35.0},
			Manufacturer:  &pb.Manufacturer{Name: "TechNova", Country: "France", Website: "https://technova.example"},
			Tags:          []string{"computer", "navigation"},
			Metadata: map[string]*pb.Value{
				"processor": {Kind: &pb.Value_StringValue{StringValue: "Quantum-7"}},
			},
			CreatedAt: ts,
			UpdatedAt: ts,
		},
		{
			Uuid:          "10",
			Name:          "Wing Gamma",
			Description:   "Auxiliary wing",
			Price:         42_000.0,
			StockQuantity: 35,
			Category:      pb.Category_WING,
			Dimensions:    &pb.Dimensions{Length: 145.0, Width: 19.0, Height: 11.0, Weight: 115.0},
			Manufacturer:  &pb.Manufacturer{Name: "EuroAero", Country: "Germany", Website: "https://euroaero.example"},
			Tags:          []string{"wing", "aux"},
			Metadata: map[string]*pb.Value{
				"material": {Kind: &pb.Value_StringValue{StringValue: "alloy"}},
			},
			CreatedAt: ts,
			UpdatedAt: ts,
		},
		// 11-15
		{
			Uuid:          "11",
			Name:          "Landing Gear A",
			Description:   "Main landing gear assembly",
			Price:         180_000.0,
			StockQuantity: 8,
			Category:      pb.Category_UNKNOWN,
			Dimensions:    &pb.Dimensions{Length: 120.0, Width: 40.0, Height: 30.0, Weight: 450.0},
			Manufacturer:  &pb.Manufacturer{Name: "RaketaTech", Country: "Russia", Website: "https://raketatech.example"},
			Tags:          []string{"landing", "gear"},
			Metadata: map[string]*pb.Value{
				"load_capacity_t": {Kind: &pb.Value_DoubleValue{DoubleValue: 50.0}},
			},
			CreatedAt: ts,
			UpdatedAt: ts,
		},
		{
			Uuid:          "12",
			Name:          "Aux Fuel Tank B2",
			Description:   "Backup fuel tank",
			Price:         60_000.0,
			StockQuantity: 12,
			Category:      pb.Category_FUEL,
			Dimensions:    &pb.Dimensions{Length: 180.0, Width: 70.0, Height: 70.0, Weight: 280.0},
			Manufacturer:  &pb.Manufacturer{Name: "Orbital Tanks", Country: "USA", Website: "https://orbitaltanks.example"},
			Tags:          []string{"fuel", "backup"},
			Metadata: map[string]*pb.Value{
				"capacity_l": {Kind: &pb.Value_DoubleValue{DoubleValue: 3500.0}},
			},
			CreatedAt: ts,
			UpdatedAt: ts,
		},
		{
			Uuid:          "13",
			Name:          "Reaction Control Thruster",
			Description:   "Small maneuvering thruster",
			Price:         95_000.0,
			StockQuantity: 20,
			Category:      pb.Category_ENGINE,
			Dimensions:    &pb.Dimensions{Length: 80.0, Width: 50.0, Height: 50.0, Weight: 150.0},
			Manufacturer:  &pb.Manufacturer{Name: "AstroWorks", Country: "USA", Website: "https://astro.example"},
			Tags:          []string{"thruster", "RCS"},
			Metadata: map[string]*pb.Value{
				"max_thrust": {Kind: &pb.Value_DoubleValue{DoubleValue: 1200.0}},
			},
			CreatedAt: ts,
			UpdatedAt: ts,
		},
		{
			Uuid:          "14",
			Name:          "Solar Panel Alpha",
			Description:   "High-efficiency solar array",
			Price:         50_000.0,
			StockQuantity: 30,
			Category:      pb.Category_UNKNOWN,
			Dimensions:    &pb.Dimensions{Length: 200.0, Width: 80.0, Height: 5.0, Weight: 100.0},
			Manufacturer:  &pb.Manufacturer{Name: "TechNova", Country: "France", Website: "https://technova.example"},
			Tags:          []string{"solar", "power"},
			Metadata: map[string]*pb.Value{
				"power_kw": {Kind: &pb.Value_DoubleValue{DoubleValue: 25.0}},
			},
			CreatedAt: ts,
			UpdatedAt: ts,
		},
		{
			Uuid:          "15",
			Name:          "Wing Delta",
			Description:   "Extra wing module",
			Price:         47_000.0,
			StockQuantity: 25,
			Category:      pb.Category_WING,
			Dimensions:    &pb.Dimensions{Length: 155.0, Width: 21.0, Height: 12.0, Weight: 125.0},
			Manufacturer:  &pb.Manufacturer{Name: "EuroAero", Country: "Germany", Website: "https://euroaero.example"},
			Tags:          []string{"wing", "extra"},
			Metadata: map[string]*pb.Value{
				"material": {Kind: &pb.Value_StringValue{StringValue: "carbon"}},
			},
			CreatedAt: ts,
			UpdatedAt: ts,
		},
		// 16-20
		{
			Uuid:          "16",
			Name:          "Porthole Glass Z1",
			Description:   "Standard porthole",
			Price:         10_000.0,
			StockQuantity: 20,
			Category:      pb.Category_PORTHOLE,
			Dimensions:    &pb.Dimensions{Length: 45.0, Width: 45.0, Height: 8.0, Weight: 45.0},
			Manufacturer:  &pb.Manufacturer{Name: "SpaceVision", Country: "Japan", Website: "https://spacevision.example"},
			Tags:          []string{"glass"},
			Metadata: map[string]*pb.Value{
				"material": {Kind: &pb.Value_StringValue{StringValue: "tempered glass"}},
			},
			CreatedAt: ts,
			UpdatedAt: ts,
		},
		{
			Uuid:          "17",
			Name:          "Auxiliary Engine X2",
			Description:   "Auxiliary engine module",
			Price:         400_000.0,
			StockQuantity: 7,
			Category:      pb.Category_ENGINE,
			Dimensions:    &pb.Dimensions{Length: 160.0, Width: 110.0, Height: 150.0, Weight: 650.0},
			Manufacturer:  &pb.Manufacturer{Name: "AstroWorks", Country: "USA", Website: "https://astro.example"},
			Tags:          []string{"engine", "aux"},
			Metadata: map[string]*pb.Value{
				"max_thrust": {Kind: &pb.Value_DoubleValue{DoubleValue: 3500.0}},
			},
			CreatedAt: ts,
			UpdatedAt: ts,
		},
	}

	for _, pbPart := range parts {
		// конвертируем protobuf → domain модель
		modelPart := PartProtoToModel(pbPart) // *model.Part

		// теперь конвертируем domain → repository модель
		repoPart := converter.ModelPartToRepo(*modelPart)

		// сохраняем в inventory (map[string]*repoModel.GetPartResponse)
		r.inventory[repoPart.UUID] = &repoModel.GetPartResponse{
			Part: repoPart, // repoModel.Part
		}
	}
}
