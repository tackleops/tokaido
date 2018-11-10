package git

// GitignoreDrupalDefaults - Return the default Drupal 8 gitignore list
func GitignoreDrupalDefaults() string {
	return `
#
# Recommended .gitinore list sourced from:
# https://github.com/WondrousLLC/drupal-8-best-practices/blob/master/.gitignore
#

# Ignore configuration files that may contain sensitive information.
sites/*/*settings*.php

# Ignore paths that contain generated content.
files/
sites/*/files
sites/*/private

# Ignore default text files
robots.txt
/CHANGELOG.txt
/COPYRIGHT.txt
/INSTALL*.txt
/LICENSE.txt
/MAINTAINERS.txt
/UPGRADE.txt
/README.txt
sites/README.txt
sites/all/modules/README.txt
sites/all/themes/README.txt

# Ignore everything but the "sites" folder ( for non core developer )
.htaccess
web.config
authorize.php
cron.php
index.php
install.php
update.php
xmlrpc.php
/includes
/misc
/modules
/profiles
/scripts
/themes

### OSX template
*.DS_Store
.AppleDouble
.LSOverride

# Icon must end with two \r
Icon

# Thumbnails
._*

# Files that might appear in the root of a volume
.DocumentRevisions-V100
.fseventsd
.Spotlight-V100
.TemporaryItems
.Trashes
.VolumeIcon.icns
.com.apple.timemachine.donotpresent

# Directories potentially created on remote AFP share
.AppleDB
.AppleDesktop
Network Trash Folder
Temporary Items
.apdisk
### JetBrains template
# Covers JetBrains IDEs: IntelliJ, RubyMine, PhpStorm, AppCode, PyCharm, CLion, Android Studio and Webstorm
# Reference: https://intellij-support.jetbrains.com/hc/en-us/articles/206544839

# User-specific stuff:
.idea
.idea/workspace.xml
.idea/tasks.xml
.idea/dictionaries
.idea/vcs.xml
.idea/jsLibraryMappings.xml

# Sensitive or high-churn files:
.idea/dataSources.ids
.idea/dataSources.xml
.idea/dataSources.local.xml
.idea/sqlDataSources.xml
.idea/dynamic.xml
.idea/uiDesigner.xml

# Gradle:
.idea/gradle.xml
.idea/libraries

# Mongo Explorer plugin:
.idea/mongoSettings.xml

## File-based project format:
*.iws

## Plugin-specific files:

# IntelliJ
/out/

# mpeltonen/sbt-idea plugin
.idea_modules/

# JIRA plugin
atlassian-ide-plugin.xml

# Crashlytics plugin (for Android Studio and IntelliJ)
com_crashlytics_export_strings.xml
crashlytics.properties
crashlytics-build.properties
fabric.properties
`
}
