#!/usr/bin/env bash

# NAME
#         openapi-generator-cli generate - Generate code with the specified
#         generator.

# SYNOPSIS
#         openapi-generator-cli generate
#                 [(-a <authorization> | --auth <authorization>)]
#                 [--api-name-suffix <api name suffix>] [--api-package <api package>]
#                 [--artifact-id <artifact id>] [--artifact-version <artifact version>]
#                 [(-c <configuration file> | --config <configuration file>)] [--dry-run]
#                 [(-e <templating engine> | --engine <templating engine>)]
#                 [--enable-post-process-file]
#                 [(-g <generator name> | --generator-name <generator name>)]
#                 [--generate-alias-as-model] [--git-host <git host>]
#                 [--git-repo-id <git repo id>] [--git-user-id <git user id>]
#                 [--global-property <global properties>...] [--group-id <group id>]
#                 [--http-user-agent <http user agent>]
#                 [(-i <spec file> | --input-spec <spec file>)]
#                 [--ignore-file-override <ignore file override location>]
#                 [--import-mappings <import mappings>...]
#                 [--inline-schema-name-mappings <inline schema name mappings>...]
#                 [--inline-schema-options <inline schema options>...]
#                 [--input-spec-root-directory <Folder with spec(s)>]
#                 [--instantiation-types <instantiation types>...]
#                 [--invoker-package <invoker package>]
#                 [--language-specific-primitives <language specific primitives>...]
#                 [--legacy-discriminator-behavior] [--library <library>]
#                 [--log-to-stderr] [--merged-spec-filename <Name of resulted merged specs file (used along with --input-spec-root-directory option)>]
#                 [--minimal-update] [--model-name-mappings <model name mappings>...]
#                 [--model-name-prefix <model name prefix>]
#                 [--model-name-suffix <model name suffix>]
#                 [--model-package <model package>]
#                 [--name-mappings <property name mappings>...]
#                 [(-o <output directory> | --output <output directory>)]
#                 [--openapi-normalizer <OpenAPI normalizer rules>...] [(-p <additional properties> | --additional-properties <additional properties>)...]
#                 [--package-name <package name>]
#                 [--parameter-name-mappings <parameter name mappings>...]
#                 [--release-note <release note>] [--remove-operation-id-prefix]
#                 [--reserved-words-mappings <reserved word mappings>...]
#                 [(-s | --skip-overwrite)] [--schema-mappings <schema mappings>...]
#                 [--server-variables <server variables>...] [--skip-operation-example]
#                 [--skip-validate-spec] [--strict-spec <true/false strict behavior>]
#                 [(-t <template directory> | --template-dir <template directory>)]
#                 [--type-mappings <type mappings>...] [(-v | --verbose)]

# OPTIONS
#         -a <authorization>, --auth <authorization>
#             adds authorization headers when fetching the OpenAPI definitions
#             remotely. Pass in a URL-encoded string of name:header with a comma
#             separating multiple values

#         --api-name-suffix <api name suffix>
#             Suffix that will be appended to all API names ('tags'). Default:
#             Api. e.g. Pet => PetApi. Note: Only ruby, python, jaxrs generators
#             support this feature at the moment.

#         --api-package <api package>
#             package for generated api classes

#         --artifact-id <artifact id>
#             artifactId in generated pom.xml. This also becomes part of the
#             generated library's filename

#         --artifact-version <artifact version>
#             artifact version in generated pom.xml. This also becomes part of the
#             generated library's filename

#         -c <configuration file>, --config <configuration file>
#             Path to configuration file. It can be JSON or YAML. If file is JSON,
#             the content should have the format {"optionKey":"optionValue",
#             "optionKey1":"optionValue1"...}. If file is YAML, the content should
#             have the format optionKey: optionValue. Supported options can be
#             different for each language. Run config-help -g {generator name}
#             command for language-specific config options.

#         --dry-run
#             Try things out and report on potential changes (without actually
#             making changes).

#         -e <templating engine>, --engine <templating engine>
#             templating engine: "mustache" (default) or "handlebars" (beta)

#         --enable-post-process-file
#             Enable post-processing file using environment variables.

#         -g <generator name>, --generator-name <generator name>
#             generator to use (see list command for list)

#         --generate-alias-as-model
#             Generate model implementation for aliases to map and array schemas.
#             An 'alias' is an array, map, or list which is defined inline in a
#             OpenAPI document and becomes a model in the generated code. A 'map'
#             schema is an object that can have undeclared properties, i.e. the
#             'additionalproperties' attribute is set on that object. An 'array'
#             schema is a list of sub schemas in a OAS document

#         --git-host <git host>
#             Git host, e.g. gitlab.com.

#         --git-repo-id <git repo id>
#             Git repo ID, e.g. openapi-generator.

#         --git-user-id <git user id>
#             Git user ID, e.g. openapitools.

#         --global-property <global properties>
#             sets specified global properties (previously called 'system
#             properties') in the format of name=value,name=value (or multiple
#             options, each with name=value)

#         --group-id <group id>
#             groupId in generated pom.xml

#         --http-user-agent <http user agent>
#             HTTP user agent, e.g. codegen_csharp_api_client, default to
#             'OpenAPI-Generator/{packageVersion}/{language}'

#         -i <spec file>, --input-spec <spec file>
#             location of the OpenAPI spec, as URL or file (required if not loaded
#             via config using -c)

#         --ignore-file-override <ignore file override location>
#             Specifies an override location for the .openapi-generator-ignore
#             file. Most useful on initial generation.

#         --import-mappings <import mappings>
#             specifies mappings between a given class and the import that should
#             be used for that class in the format of type=import,type=import. You
#             can also have multiple occurrences of this option.

#         --inline-schema-name-mappings <inline schema name mappings>
#             specifies mappings between the inline schema name and the new name
#             in the format of inline_object_2=Cat,inline_object_5=Bird. You can
#             also have multiple occurrences of this option.

#         --inline-schema-options <inline schema options>
#             specifies the options for handling inline schemas in the inline
#             model resolver. Please refer to https://github.com/OpenAPITools/openapi-generator/blob/master/docs/customization.md
#             for a list of options.

#         --input-spec-root-directory <Folder with spec(s)>
#             Local root folder with spec file(s)

#         --instantiation-types <instantiation types>
#             sets instantiation type mappings in the format of
#             type=instantiatedType,type=instantiatedType.For example (in Java):
#             array=ArrayList,map=HashMap. In other words array types will get
#             instantiated as ArrayList in generated code. You can also have
#             multiple occurrences of this option.

#         --invoker-package <invoker package>
#             root package for generated code

#         --language-specific-primitives <language specific primitives>
#             specifies additional language specific primitive types in the format
#             of type1,type2,type3,type3. For example:
#             String,boolean,Boolean,Double. You can also have multiple
#             occurrences of this option.

#         --legacy-discriminator-behavior
#             Set to false for generators with better support for discriminators.
#             (Python, Java, Go, PowerShell, C# have this enabled by default).

#         --library <library>
#             library template (sub-template)

#         --log-to-stderr
#             write all log messages (not just errors) to STDOUT. Useful for
#             piping the JSON output of debug options (e.g. `--global-property
#             debugOperations`) to an external parser directly while testing a
#             generator.

#         --merged-spec-filename <Name of resulted merged specs file (used along
#         with --input-spec-root-directory option)>

#         --minimal-update
#             Only write output files that have changed.

#         --model-name-mappings <model name mappings>
#             specifies mappings between the model name and the new name in the
#             format of model_name=AnotherName,model_name2=OtherName2. You can
#             also have multiple occurrences of this option.

#         --model-name-prefix <model name prefix>
#             Prefix that will be prepended to all model names.

#         --model-name-suffix <model name suffix>
#             Suffix that will be appended to all model names.

#         --model-package <model package>
#             package for generated models

#         --name-mappings <property name mappings>
#             specifies mappings between the property name and the new name in the
#             format of prop_name=PropName,prop_name2=PropName2. You can also have
#             multiple occurrences of this option.

#         -o <output directory>, --output <output directory>
#             where to write the generated files (current dir by default)

#         --openapi-normalizer <OpenAPI normalizer rules>
#             specifies the rules to be enabled in OpenAPI normalizer in the form
#             of RULE_1=true,RULE_2=original. You can also have multiple
#             occurrences of this option.

#         -p <additional properties>, --additional-properties <additional
#         properties>
#             sets additional properties that can be referenced by the mustache
#             templates in the format of name=value,name=value. You can also have
#             multiple occurrences of this option.

#         --package-name <package name>
#             package for generated classes (where supported)

#         --parameter-name-mappings <parameter name mappings>
#             specifies mappings between the parameter name and the new name in
#             the format of param_name=paramName,param_name2=paramName2. You can
#             also have multiple occurrences of this option.

#         --release-note <release note>
#             Release note, default to 'Minor update'.

#         --remove-operation-id-prefix
#             Remove prefix of operationId, e.g. config_getId => getId

#         --reserved-words-mappings <reserved word mappings>
#             specifies how a reserved name should be escaped to. Otherwise, the
#             default _<name> is used. For example id=identifier. You can also
#             have multiple occurrences of this option.

#         -s, --skip-overwrite
#             specifies if the existing files should be overwritten during the
#             generation.

#         --schema-mappings <schema mappings>
#             specifies mappings between the schema and the new name in the format
#             of schema_a=Cat,schema_b=Bird. You can also have multiple
#             occurrences of this option.

#         --server-variables <server variables>
#             sets server variables overrides for spec documents which support
#             variable templating of servers.

#         --skip-operation-example
#             Skip examples defined in operations to avoid out of memory errors.

#         --skip-validate-spec
#             Skips the default behavior of validating an input specification.

#         --strict-spec <true/false strict behavior>
#             'MUST' and 'SHALL' wording in OpenAPI spec is strictly adhered to.
#             e.g. when false, no fixes will be applied to documents which pass
#             validation but don't follow the spec.

#         -t <template directory>, --template-dir <template directory>
#             folder containing the template files

#         --type-mappings <type mappings>
#             sets mappings between OpenAPI spec types and generated code types in
#             the format of OpenAPIType=generatedType,OpenAPIType=generatedType.
#             For example: array=List,map=Map,string=String. You can also have
#             multiple occurrences of this option. To map a specified format, use
#             type+format, e.g. string+password=EncryptedString will map `type:
#             string, format: password` to `EncryptedString`.

#         -v, --verbose
#             verbose mode

# ---

# ❯ npx @openapitools/openapi-generator-cli config-help -g go

# CONFIG OPTIONS

# 	disallowAdditionalPropertiesIfNotPresent
# 	    If false, the 'additionalProperties' implementation (set to true by default) is compliant with the OAS and JSON schema specifications. If true (default), keep the old (incorrect) behaviour that 'additionalProperties' is set to false by default. (Default: true)
# 	        false - The 'additionalProperties' implementation is compliant with the OAS and JSON schema specifications.
# 	        true - Keep the old (incorrect) behaviour that 'additionalProperties' is set to false by default.

# 	enumClassPrefix
# 	    Prefix enum with class name (Default: false)

# 	generateInterfaces
# 	    Generate interfaces for api classes (Default: false)

# 	hideGenerationTimestamp
# 	    Hides the generation timestamp when files are generated. (Default: true)

# 	isGoSubmodule
# 	    whether the generated Go module is a submodule (Default: false)

# 	packageName
# 	    Go package name (convention: lowercase). (Default: openapi)

# 	packageVersion
# 	    Go package version. (Default: 1.0.0)

# 	prependFormOrBodyParameters
# 	    Add form or body parameters to the beginning of the parameter list. (Default: false)

# 	structPrefix
# 	    whether to prefix struct with the class name. e.g. DeletePetOpts => PetApiDeletePetOpts (Default: false)

# 	useOneOfDiscriminatorLookup
# 	    Use the discriminator's mapping in oneOf to speed up the model lookup. IMPORTANT: Validation (e.g. one and only one match in oneOf's schemas) will be skipped. (Default: false)

# 	withAWSV4Signature
# 	    whether to include AWS v4 signature support (Default: false)

# 	withXml
# 	    whether to include support for application/xml content type and include XML annotations in the model (works with libraries that provide support for JSON and XML) (Default: false)

GITHUB_USER_ID=vlazic
GITHUB_REPO_ID=opalstack-go
BASE_PATH=https://my.opalstack.com/api/v1

additional_properties=(
    "packageName=opalstack"
    "packageVersion=0.0.1"
    "basePath=$BASE_PATH"
    "apiKeyPrefix=Bearer"
    "withGoCodegenComment=true"
    "goModuleName=github.com/$GITHUB_USER_ID/$GITHUB_REPO_ID"
)

additional_properties_string=$(
    IFS=,
    echo "${additional_properties[*]}"
)

npx @openapitools/openapi-generator-cli generate \
    -i https://my.opalstack.com/api/v1/schema/ \
    -g go \
    -o . \
    --package-name OpalStack \
    --git-user-id $GITHUB_USER_ID \
    --git-repo-id $GITHUB_REPO_ID \
    --additional-properties="$additional_properties_string"
