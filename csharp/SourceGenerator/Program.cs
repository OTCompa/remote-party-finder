using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Runtime.CompilerServices;
using System.Text;
using Lumina;
using Lumina.Data;
using Lumina.Excel;
using Lumina.Excel.Sheets;
using Lumina.Text;
using Lumina.Text.Parse;
using Lumina.Text.ReadOnly;
using Pidgin;

namespace SourceGenerator;

internal class Program
{
    private static void Main(string[] args)
    {
        if (args.Length < 2)
        {
            Console.WriteLine($"Usage: SourceGenerator <sqpack dir> <out dir>");
            return;
        }

        var data = new Dictionary<Language, GameData>(4);
        foreach (var lang in Languages.Keys)
        {
            data[lang] = new GameData(args[0], new LuminaOptions
            {
                PanicOnSheetChecksumMismatch = false,
                DefaultExcelLanguage = lang,
            });
        }

        var prog = new Program(data);
        var outPath = args[1];
        File.WriteAllText(Path.Join(outPath, "auto_translate.go"), prog.GenerateAutoTranslate());
        File.WriteAllText(Path.Join(outPath, "content_kinds.go"), prog.GenerateContentKinds());
        File.WriteAllText(Path.Join(outPath, "duties.go"), prog.GenerateDuties());
        File.WriteAllText(Path.Join(outPath, "jobs.go"), prog.GenerateJobs());
        File.WriteAllText(Path.Join(outPath, "roulettes.go"), prog.GenerateRoulettes());
        File.WriteAllText(Path.Join(outPath, "territory_names.go"), prog.GenerateTerritoryNames());
        File.WriteAllText(Path.Join(outPath, "treasure_maps.go"), prog.GenerateTreasureMaps());
        File.WriteAllText(Path.Join(outPath, "worlds.go"), prog.GenerateWorlds());
    }

    private Dictionary<Language, GameData> Data { get; }

    private Program(Dictionary<Language, GameData> data)
    {
        this.Data = data;
    }

    private static StringBuilder DefaultHeader(bool localisedText = false)
    {
        var sb = new StringBuilder("package ffxiv\n");

        /*
        if (localisedText)
        {
            sb.Append("use super::LocalisedText;\n");
        }
        */

        return sb;
    }

    private static readonly Dictionary<Language, string> Languages = new()
    {
        [Language.English] = "En",
        [Language.Japanese] = "Ja",
        [Language.German] = "De",
        [Language.French] = "Fr",
    };

    private string? GetLocalisedStruct<T>(uint rowId, Func<T, ReadOnlySeString?> nameFunc, uint indent = 0,
        bool capitalise = false, bool type = true) where T : struct, IExcelRow<T>
    {
        var def = this.Data[Language.English].GetExcelSheet<T>()!.GetRow(rowId)!;
        var defName = nameFunc(def)?.ExtractText();
        if (string.IsNullOrEmpty(defName))
        {
            return null;
        }

        var sb = new StringBuilder();
        if (type)
        {
            sb.Append("LocalisedText{\n");
        }
        else
        {
            sb.Append("{\n");
        }

        foreach (var (language, key) in Languages)
        {
            var row = this.Data[language].GetExcelSheet<T>(language)?.GetRow(rowId);
            var name = row == null
                ? defName
                : nameFunc((T)row)?.ExtractText().Replace("\"", "\\\"").Replace(" ", "").Replace("­", "");
            name ??= defName;
            if (capitalise)
            {
                name = name[..1].ToUpperInvariant() + name[1..];
            }

            for (var i = 0; i < indent + 2; i++)
            {
                sb.Append(' ');
            }

            sb.Append($"{key}: \"{name}\",\n");
        }

        for (var i = 0; i < indent; i++)
        {
            sb.Append(' ');
        }

        sb.Append('}');

        return sb.ToString();
    }

    private string GenerateDuties()
    {
        var sb = DefaultHeader(true);
        sb.Append('\n');

        sb.Append("type DutyInfo struct {\n");
        sb.Append("\tName        LocalisedText\n");
        sb.Append("\tHighEnd     bool\n");
        sb.Append("\tContentKind ContentKind\n");
        sb.Append("}\n\n");

        /*
        sb.Append("#[derive(Debug, Clone, Copy)]\n");
        sb.Append("#[allow(unused)]\n");
        sb.Append("#[repr(u32)]\n");
        sb.Append("pub enum ContentKind {\n");
        foreach (var kind in this.Data[Language.English].GetExcelSheet<ContentType>()!)
        {
            var name = kind.Name.ExtractText().Replace(" ", "").Replace("&", "");
            if (name.Length > 0)
            {
                sb.Append($"    {name} = {kind.RowId},\n");
            }
        }

        sb.Append("    Other(u32),\n");
        sb.Append("}\n\n");

        sb.Append("impl ContentKind {\n");

        sb.Append("    fn from_u32(kind: u32) -> Self {\n");
        sb.Append("        match kind {\n");
        foreach (var kind in this.Data[Language.English].GetExcelSheet<ContentType>()!)
        {
            var name = kind.Name.ExtractText().Replace(" ", "").Replace("&", "");
            if (name.Length > 0)
            {
                sb.Append($"            {kind.RowId} => Self::{name},\n");
            }
        }

        sb.Append("            x => Self::Other(x),\n");
        sb.Append("        }\n");
        sb.Append("    }\n\n");

        sb.Append("    pub fn as_u32(self) -> u32 {\n");
        sb.Append("        match self {\n");
        foreach (var kind in this.Data[Language.English].GetExcelSheet<ContentType>()!)
        {
            var name = kind.Name.ExtractText().Replace(" ", "").Replace("&", "");
            if (name.Length > 0)
            {
                sb.Append($"            Self::{name} => {kind.RowId},\n");
            }
        }

        sb.Append("            Self::Other(x) => x,\n");
        sb.Append("        }\n");
        sb.Append("    }\n");

        sb.Append("}\n\n");

        sb.Append("lazy_static::lazy_static! {\n");
        */
        sb.Append("var DUTIES = map[uint32]DutyInfo{\n");

        foreach (var cfc in this.Data[Language.English].GetExcelSheet<ContentFinderCondition>()!)
        {
            if (cfc.RowId == 0)
            {
                continue;
            }

            var name = this.GetLocalisedStruct<ContentFinderCondition>(cfc.RowId, row => row.Name, 12, true);
            if (name == null)
            {
                continue;
            }

            var highEnd = cfc.HighEndDuty ? "true" : "false";
            var contentType = cfc.ContentType.Value;
            var contentKind = contentType.Name.ExtractText().Replace(" ", "").Replace("&", "");
            if (string.IsNullOrEmpty(contentKind))
            {
                contentKind = $"({contentType.RowId})";
            }

            sb.Append($"\t{cfc.RowId}: {{\n");
            sb.Append($"\t\tName: {name},\n");
            sb.Append($"\t\tHighEnd: {highEnd},\n");
            sb.Append($"\t\tContentKind: ContentKind{contentKind},\n");
            sb.Append("\t},\n");
        }

        sb.Append("}\n");

        return sb.ToString();
    }

    private string GenerateJobs()
    {
        var sb = DefaultHeader();
        sb.Append("type ClassJob interface {\n");
        sb.Append("\tisClassJob()\n");
        sb.Append("}\n\n");

        sb.Append("type Class string\n\n");
        sb.Append("func (Class) isClassJob() {}\n\n");

        sb.Append("type Job string\n\n");
        sb.Append("func (Job) isClassJob() {}\n\n");

        sb.Append("type NonCombatJob string\n\n");
        sb.Append("func (NonCombatJob) isClassJob() {}\n\n");

        List<string> classes = [];
        List<string> jobs = [];
        List<string> nonCombatJobs = [];
        foreach (var cj in this.Data[Language.English].GetExcelSheet<ClassJob>()!)
        {
            if (cj.RowId == 0)
            {
                continue;
            }

            var name = cj.NameEnglish.ExtractText().Replace(" ", "");
            if (name.Length <= 0)
            {
                continue;
            }

            var isCombat = cj.Role != 0;
            var isClass = cj.JobIndex == 0;
            if (!isCombat)
            {
                nonCombatJobs.Add(name);
                continue;
            }

            if (isClass)
            {
                classes.Add(name);
                continue;
            }

            jobs.Add(name);
        }

        void constGenerator(List<string> classJob, string type)
        {
            sb.Append("const (\n");
            foreach(var name in classJob)
            {
                sb.Append($"\t{name}\t{type} = \"{name}\"\n");
            }
            sb.Append(")\n\n");
        }

        constGenerator(classes, "Class");
        constGenerator(jobs, "Job");
        constGenerator(nonCombatJobs, "NonCombatJob");

        sb.Append("var JOBS = map[uint32]ClassJob{\n");
        foreach (var cj in this.Data[Language.English].GetExcelSheet<ClassJob>()!)
        {
            if (cj.RowId == 0)
            {
                continue;
            }

            var name = cj.NameEnglish.ExtractText().Replace(" ", "");
            if (name.Length <= 0)
            {
                continue;
            }

            var isCombat = cj.Role != 0;
            var isClass = cj.JobIndex == 0;

            sb.Append($"\t{cj.RowId}: {name},\n");
        }

        sb.Append("}\n");

        return sb.ToString();
    }

    private string GenerateRoulettes()
    {
        var sb = DefaultHeader(true);
        sb.Append('\n');
        sb.Append("type RouletteInfo struct {\n");
        sb.Append("\tName LocalisedText\n");
        sb.Append("\tPVP  bool\n");
        sb.Append("}\n\n");

        sb.Append("var ROULETTES = map[uint32]RouletteInfo{\n");

        foreach (var cr in this.Data[Language.English].GetExcelSheet<ContentRoulette>()!)
        {
            if (cr.RowId == 0)
            {
                continue;
            }

            var name = this.GetLocalisedStruct<ContentRoulette>(cr.RowId, row => row.Name, 12);
            if (name == null)
            {
                continue;
            }

            var pvp = cr.IsPvP
                ? "true"
                : "false";

            sb.Append($"\t{cr.RowId}: {{\n");
            sb.Append($"\t\tName: {name},\n");
            sb.Append($"\t\tPVP: {pvp},\n");
            sb.Append("\t},\n");
        }
        sb.Append("}\n");

        return sb.ToString();
    }

    private string GenerateWorlds()
    {
        var sb = DefaultHeader();
        var datacenters = new Dictionary<string, List<string>>();
        string[] regions = ["JP", "NA", "EU", "OC", "Others"];

        // datacenters
        sb.Append("type Datacenter struct {\n");
        sb.Append("\tName string\n");
        sb.Append("\tId int\n");
        sb.Append("}\n\n");

        sb.Append("var (\n");

        foreach (var region in regions)
        {
            datacenters[region] = [];
        }

        foreach (var dc in this.Data[Language.English].GetExcelSheet<WorldDCGroupType>()!)
        {
            if (dc.RowId == 0) { continue; }
            var name = dc.Name.ExtractText();

            // remove non JP/NA/EU/OCE dcs
            if ((int)dc.Region < 1 || (int)dc.Region > 4)
            { 
                continue; 
            }
            sb.Append($"\t{name} Datacenter = Datacenter{{Name:\"{name}\", Id : {dc.NeolobbyId}}}\n");

            // add datacenter regions
            switch ((int)dc.Region) {
                case 1:
                    datacenters[regions[0]].Add(name); break;
                case 2:
                    datacenters[regions[1]].Add(name); break;
                case 3:
                    datacenters[regions[2]].Add(name); break;
                case 4:
                    datacenters[regions[3]].Add(name); break;
                default:
                    datacenters[regions[4]].Add(name); break;
            }
        }
        sb.Append(")\n\n");

        // add datacenter regions
        foreach (var region in regions)
        {
            sb.Append($"func Datacenter{region}() []Datacenter {{\n");
            sb.Append("\treturn []Datacenter{");
            foreach (var world in datacenters[region])
            {
                sb.Append(world + ",");
            }
            sb.Append("}\n}\n\n");
        }

        // world struct
        sb.Append("type World struct {\n");
        sb.Append("\tName       string\n");
        sb.Append("\tDatacenter Datacenter\n");
        sb.Append("}\n\n");

        // world constants
        sb.Append("var (\n");
        foreach (var world in this.Data[Language.English].GetExcelSheet<World>()!)
        {
            if (world.RowId == 0 || !world.IsPublic || world.UserType == 0 || world.DataCenter.RowId == 0)
            {
                continue;
            }

            var name = world.Name.ExtractText();
            if (name.Length <= 0) 
            {
                continue;
            }

            // remove cloud dcs
            if (name.Contains("0"))
            {
                continue;
            }

            var dc = world.DataCenter.Value.Name.ExtractText().Replace(" ", "").Replace("(Beta)", "");
            sb.Append($"\t{name}\tWorld = World{{Name: \"{name}\", Datacenter: {dc}}}\n");
        }
        sb.Append(")\n\n");

        sb.Append("var WORLDS = map[uint32]World{\n");
        foreach (var world in this.Data[Language.English].GetExcelSheet<World>()!)
        {
            if (world.RowId == 0 || !world.IsPublic || world.UserType == 0 || world.DataCenter.RowId == 0)
            {
                continue;
            }

            var name = world.Name.ExtractText();
            if (name.Length <= 0)
            {
                continue;
            }

            // remove cloud dcs
            if (name.Contains("0"))
            {
                continue;
            }

            sb.Append($"\t{world.RowId}: {name},\n");
        }

        sb.Append("}\n");

        return sb.ToString();
    }

    private string GenerateTerritoryNames()
    {
        var sb = DefaultHeader(true);
        sb.Append("var TERRITORY_NAMES = map[uint32]LocalisedText{\n");

        foreach (var tt in this.Data[Language.English].GetExcelSheet<TerritoryType>()!)
        {
            if (tt.RowId == 0 || tt.PlaceName.RowId == 0)
            {
                continue;
            }

            var name = this.GetLocalisedStruct<TerritoryType>(
                tt.RowId,
                row => row.PlaceName.Value!.Name,
                8,
                false,
                false
            );
            if (name == null)
            {
                continue;
            }

            sb.Append($"\t{tt.RowId}: {name},\n");
        }

        sb.Append("}\n");

        return sb.ToString();
    }

    private string GenerateAutoTranslate()
    {
        var sb = DefaultHeader(true);

        sb.Append("\ntype MapStruct struct {\n");
        sb.Append("\tRowGroup, RowId uint32\n");
        sb.Append("}\n\n");


        sb.Append("var AUTO_TRANSLATE = map[MapStruct]LocalisedText{\n");

        var parser = AutoTranslate.Parser();
        foreach (var row in this.Data[Language.English].GetExcelSheet<Completion>()!)
        {
            var lookup = row.LookupTable.ToString().Replace("<num(", "").Replace(")>", ""); // 🙂
            
            if (lookup is not ("" or "@"))
            {
                string sheetName;
                Maybe<IEnumerable<ISelectorPart>> selector;

                try
                {
                    (sheetName, selector) = parser.ParseOrThrow(lookup);
                }
                catch (Exception ex)
                {
                    Console.WriteLine(lookup);
                    Console.WriteLine(ex.Message);
                    continue;
                }

                var sheets = this.Data
                    .Select(kv => (kv.Key, new ExcelSheet<RawRow>(kv.Value.Excel.GetRawSheet(sheetName))))
                    .ToDictionary();

                var columns = new List<int>();
                var rows = new List<Range>();
                if (selector.HasValue)
                {
                    columns.Clear();
                    rows.Clear();

                    foreach (var part in selector.Value)
                    {
                        switch (part)
                        {
                            case IndexRange range:
                            {
                                var start = (int)range.Start;
                                var end = (int)(range.End + 1);
                                rows.Add(start..end);
                                break;
                            }
                            case SingleRow single:
                            {
                                var idx = (int)single.Row;
                                rows.Add(idx..(idx + 1));
                                break;
                            }
                            case ColumnSpecifier col:
                                columns.Add((int)col.Column);
                                break;
                        }
                    }
                }

                if (columns.Count == 0)
                {
                    columns.Add(0);
                }

                if (rows.Count == 0)
                {
                    rows.Add(..);
                }

                var builder = new StringBuilder();
                foreach (var range in rows)
                {
                    for (var i = (uint)range.Start.Value; i < range.End.Value; i++)
                    {
                        if (!sheets[Language.English].HasRow(i))
                            continue;

                        builder.Clear();
                        builder.Append($"\t{{{row.Group}, {i}}}: {{\n");

                        var lines = 0;
                        foreach (var lang in this.Data.Keys)
                        {
                            var sheet = sheets[lang];

                            var idx = i;
                            foreach (var text in from col in columns
                                     let rawRow = sheet.GetRow(idx)
                                     select rawRow.ReadStringColumn(col).ExtractText()
                                     into text
                                     where text.Length > 0
                                     select text)
                            {
                                var replace = text.Replace(" ", "").Replace("­", "").Replace("\n", "");

                                builder.Append($"            {Languages[lang]}: \"{replace}\",\n");
                                lines += 1;
                                break;
                            }
                        }

                        builder.Append("\t\t},\n");

                        if (lines != 4)
                        {
                            continue;
                        }

                        sb.Append(builder);
                    }
                }
            }
            else
            {
                var text = this.GetLocalisedStruct<Completion>(row.RowId, row => row.Text, 8, false, false);
                if (text != null)
                {
                    sb.Append($"\t{{{row.Group}, {row.RowId}}}: {text},\n");
                }
            }
        }

        sb.Append("}\n");

        return sb.ToString();
    }

    private string GenerateTreasureMaps()
    {
        var sb = DefaultHeader(true);
        sb.Append("\nvar TREASURE_MAPS = map[uint32]LocalisedText{\n");
        sb.Append("\t0: {\n");
        sb.Append("\t\tEn: \"All Levels\",\n");
        sb.Append("\t\tJa: \"レベルを指定しない\",\n");
        sb.Append("\t\tDe: \"Jede Stufe\",\n");
        sb.Append("\t\tFr: \"Tous niveaux\",\n");
        sb.Append("\t},\n");

        var i = 1;
        foreach (var row in this.Data[Language.English].GetExcelSheet<TreasureHuntRank>()!)
        {
            // IS THIS RIGHT?
            if (row.Icon == 0 || row.TreasureHuntTexture != 0)
            {
                continue;
            }

            ReadOnlySeString? GetMapName(TreasureHuntRank thr)
            {
                var name = thr.KeyItemName.Value.Name;
                return string.IsNullOrEmpty(name.ExtractText())
                    ? thr.ItemName.Value.Name
                    : name;
            }

            var name = this.GetLocalisedStruct<TreasureHuntRank>(row.RowId, GetMapName, 8, false, false);
            if (!string.IsNullOrEmpty(name))
            {
                sb.Append($"\t{i++}: {name},\n");
            }
        }

        sb.Append("}\n");

        return sb.ToString();
    }

    private string GenerateContentKinds()
    {
        var sb = DefaultHeader(true);
        sb.Append("type ContentKind uint32\n\n");

        // consts
        bool isFirst = true;
        int ctr = 1;
        sb.Append("const (\n");
        foreach (var kind in this.Data[Language.English].GetExcelSheet<ContentType>()!)
        {
            var name = kind.Name.ExtractText().Replace(" ", "").Replace("&", "");
            if (name.Length > 0)
            {
                while (ctr != kind.RowId && ctr < 100)
                {
                    sb.Append($"\t_ // {ctr}\n");
                    ctr++;
                }

                sb.Append($"\tContentKind{name}");
                if (isFirst)
                {
                    sb.Append(" ContentKind = 1 + iota\n");
                    isFirst = false;
                } else
                {
                    sb.Append("\n");
                }
                ctr++;
            }
        }

        sb.Append("\tContentKindOther\n");
        sb.Append(")\n\n");

        sb.Append("func FromUint32(kind uint32) ContentKind {\n");
        sb.Append($"\tif kind >= 1 && kind < {ctr} {{\n");
        sb.Append("\t\treturn ContentKind(kind)\n");
        sb.Append("\t}\n");
        sb.Append("\treturn ContentKindOther\n");
        sb.Append("}\n\n");


        sb.Append("func (k ContentKind) AsUint32() uint32 {\n\treturn uint32(k)\n}\n\n");


        sb.Append("func (k ContentKind) String() string {\n");
        sb.Append("\tswitch k {\n");
        foreach (var kind in this.Data[Language.English].GetExcelSheet<ContentType>()!)
        {
            var name = kind.Name.ExtractText().Replace(" ", "").Replace("&", "");
            if (name.Length > 0)
            {
                sb.Append($"\tcase ContentKind{name}:\n");
                sb.Append($"\t\treturn \"{name}\"\n");
            }
        }

        sb.Append("\tdefault:\n");
        sb.Append("\t\treturn \"Other\"\n");
        sb.Append("\t}\n");
        sb.Append("}\n");

        return sb.ToString();
    }
}