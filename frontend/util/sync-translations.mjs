import fs from "fs/promises";
import path from "path";
import yaml from "js-yaml";
import xliff from "xliff";

async function syncXliffFiles() {
  try {
    const xliffDir = "./xliff";
    const translationsDir = "../translations";

    // Read all files from xliff directory
    const xliffFiles = await fs.readdir(xliffDir);
    const xliffFilesFiltered = xliffFiles.filter((file) =>
      file.endsWith(".xlf"),
    );

    for (const xliffFile of xliffFilesFiltered) {
      await processXliffFile(xliffFile, xliffDir, translationsDir);
    }
  } catch (error) {
    console.error("Error processing XLIFF files:", error);
  }
}

async function processXliffFile(xliffFileName, xliffDir, translationsDir) {
  try {
    // Read XLIFF file
    const xliffPath = path.join(xliffDir, xliffFileName);
    const xliffContent = await fs.readFile(xliffPath, "utf8");

    // Parse XLIFF using xliff library
    const xliffData = await xliff.xliff12ToJs(xliffContent);

    // Generate corresponding YAML filename (replace - with _)
    const baseName = path.parse(xliffFileName).name;
    const yamlFileName = baseName.replace(/-/g, "_") + ".yaml";
    const yamlPath = path.join(translationsDir, yamlFileName);

    // Check if YAML file exists
    let yamlTranslations = {};
    try {
      const yamlContent = await fs.readFile(yamlPath, "utf8");
      yamlTranslations = (yaml.load(yamlContent) || {})[baseName];
    } catch (error) {
      return;
    }

    // Update XLIFF with YAML translations
    let updatedCount = 0;

    // Navigate through the XLIFF structure
    const resources = xliffData.resources;
    for (const resource of Object.values(resources)) {
      for (const translation of Object.entries(resource)) {
        const [id, transUnit] = translation;

        // Update only if target is missing
        if (!transUnit.target) {
          // Look up translation in YAML using dot notation
          const translation = getNestedValue(yamlTranslations, id);
          if (translation && typeof translation === "string") {
            transUnit.target = translation;
            updatedCount++;
          }
        }
      }
    }

    // Convert back to XLIFF format and save
    if (updatedCount > 0) {
      const updatedXliffContent = await xliff.jsToXliff12(xliffData);
      await fs.writeFile(xliffPath, updatedXliffContent, "utf8");
      console.log(
        `Updated ${xliffFileName} with ${updatedCount} new translations.`,
      );
    }
  } catch (error) {
    console.error(`Error processing ${xliffFileName}:`, error);
  }
}

function getNestedValue(obj, path) {
  return path.split(".").reduce((current, key) => {
    return current && current[key] !== undefined ? current[key] : undefined;
  }, obj);
}

// Run the script
syncXliffFiles();

export { syncXliffFiles };
